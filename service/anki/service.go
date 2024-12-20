package anki

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/types"
	"github.com/MatthewAraujo/anki_ia/utils"
	"github.com/openai/openai-go"

	"github.com/ledongthuc/pdf"
)

type Service struct {
	db *repository.Queries

	dbTx         *sql.DB
	openAiClient *openai.Client
}

func NewService(db *repository.Queries, dbTx *sql.DB, openAiClient *openai.Client) *Service {
	return &Service{
		db:           db,
		dbTx:         dbTx,
		openAiClient: openAiClient,
	}
}

func (s *Service) BeginTransaction(ctx context.Context) (*repository.Queries, *sql.Tx, error) {
	tx, err := s.dbTx.BeginTx(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateAnki(payload *types.CreateAnkiPayload) (types.CreateAnkiResponse, int, error) {
	logger.Info("Processando o arquivo do payload")

	ctx := context.Background()
	txQueries, tx, err := s.BeginTransaction(ctx)
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	pdf, err := txQueries.CreatePdf(ctx, repository.CreatePdfParams{
		UserID:      payload.UserID,
		Filename:    payload.Name,
		TextContent: utils.ToNullString(""),
	})

	tempFile, err := os.CreateTemp("", "uploaded_*.pdf")
	if err != nil {
		err = txQueries.UpdateStatus(ctx, repository.UpdateStatusParams{
			Status: utils.ToNullString("failed"),
			ID:     pdf.ID,
		})
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, payload.File)
	if err != nil {
		err = txQueries.UpdateStatus(ctx, repository.UpdateStatusParams{
			Status: utils.ToNullString("failed"),
			ID:     pdf.ID,
		})
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		err = txQueries.UpdateStatus(ctx, repository.UpdateStatusParams{
			Status: utils.ToNullString("failed"),
			ID:     pdf.ID,
		})
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao abrir arquivo temporário: %w", err)
	}
	defer file.Close()

	text, err := extractTextFromPDF(file)
	if err != nil {
		err = txQueries.UpdateStatus(ctx, repository.UpdateStatusParams{
			Status: utils.ToNullString("failed"),
			ID:     pdf.ID,
		})
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao extrair texto do PDF: %w", err)
	}
	err = txQueries.UpdateStatusAndText(ctx, repository.UpdateStatusAndTextParams{
		Status:      utils.ToNullString("processed"),
		TextContent: utils.ToNullString(text),
		ID:          pdf.ID,
	})

	// chatCompletion, err := s.openAiClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
	// 	Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
	// 		openai.AssistantMessage(utils.GetPrompt()),
	// 		openai.UserMessage(text),
	// 	}),
	// 	Model: openai.F(openai.ChatModelGPT4o),
	// })
	// if err != nil {
	// 	panic(err.Error())
	// }
	// logger.Info(chatCompletion.Choices[0].Message.Content)

	answerOpenAi := utils.GetAnswer()

	// Criar estrutura de questions a partir da resposta do OpenAI
	questions, err := utils.ParseQuestionsFromOpenAi(answerOpenAi)
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("parsing questions from openai %w", err)
	}

	err = addQuestionsAndOptionsToDB(txQueries, pdf.ID, questions)
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("error adding questions and options to db %w", err)
	}

	questionsPDF, err := txQueries.GetQuestionsByPdfId(ctx, pdf.ID)
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("error fetching questions: %w", err)
	}

	var anki types.Anki

	for _, question := range questionsPDF {
		logger.Info("Criando resposta para o usuário")
		options, err := txQueries.GetOptionsByQuestionId(ctx, question.ID)
		if err != nil {
			return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("error fetching options for question %d: %w", question.ID, err)
		}

		alternatives := make(map[string]string)
		var rightAnswer string
		for _, option := range options {
			alternatives[option.OptionKey] = option.OptionText
			if option.IsCorrect {
				rightAnswer = option.OptionKey
			}
		}

		anki.Question = append(anki.Question, types.Question{
			ID:           question.ID,
			Question:     question.QuestionText,
			Alternatives: alternatives,
			Right_answer: rightAnswer,
		})
	}

	return types.CreateAnkiResponse{
		Question: anki.Question,
	}, http.StatusCreated, nil
}

func (s *Service) GetAnkiById(payload *types.GetAnkiByIdPayload) (types.GetAnkiByIdResponse, int, error) {
	var response types.GetAnkiByIdResponse

	_, err := s.db.GetPdfById(context.Background(), payload.Id)
	if err != nil {
		return response, http.StatusInternalServerError, fmt.Errorf("error fetching PDF: %w", err)
	}

	questions, err := s.db.GetQuestionsByPdfId(context.Background(), payload.Id)
	if err != nil {
		return response, http.StatusInternalServerError, fmt.Errorf("error fetching questions: %w", err)
	}

	var anki types.Anki

	for _, question := range questions {
		options, err := s.db.GetOptionsByQuestionId(context.Background(), question.ID)
		if err != nil {
			return response, http.StatusInternalServerError, fmt.Errorf("error fetching options for question %d: %w", question.ID, err)
		}

		alternatives := make(map[string]string)
		var rightAnswer string
		for _, option := range options {
			alternatives[option.OptionKey] = option.OptionText
			if option.IsCorrect {
				rightAnswer = option.OptionKey
			}
		}

		anki.Question = append(anki.Question, types.Question{
			ID:           question.ID,
			Question:     question.QuestionText,
			Alternatives: alternatives,
			Right_answer: rightAnswer,
		})
	}

	response.Anki = anki
	return response, http.StatusOK, nil
}

func extractTextFromPDF(file *os.File) (string, error) {
	size, err := file.Stat()
	reader, err := pdf.NewReader(file, size.Size())
	if err != nil {
		return "", fmt.Errorf("erro ao criar reader para o PDF: %w", err)
	}

	var buf bytes.Buffer
	for pageIndex := 1; pageIndex <= reader.NumPage(); pageIndex++ {
		page := reader.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}
		text, _ := page.GetPlainText(nil)
		buf.WriteString(text)
	}
	return buf.String(), nil
}

func addQuestionsAndOptionsToDB(txQueries *repository.Queries, pdfID int32, questions []types.Question) error {
	for _, question := range questions {
		questionID, err := txQueries.CreateQuestion(context.Background(), repository.CreateQuestionParams{
			PdfID:        pdfID,
			QuestionText: question.Question,
		})
		if err != nil {
			return fmt.Errorf("error inserting question: %w", err)
		}

		for key, optionText := range question.Alternatives {
			isCorrect := false
			if key == question.Right_answer {
				isCorrect = true
			}

			err := txQueries.InsertOption(context.Background(), repository.InsertOptionParams{
				QuestionID: questionID,
				OptionKey:  key,
				OptionText: optionText,
				IsCorrect:  isCorrect,
			})
			if err != nil {
				return fmt.Errorf("error inserting options: %w", err)
			}
		}
	}
	return nil
}
