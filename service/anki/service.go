package anki

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/types"
	"github.com/MatthewAraujo/anki_ia/utils"
	"github.com/google/uuid"
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

	defer tx.Rollback()

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateAnki(payload *types.CreateAnkiPayload) (types.CreateAnkiResponse, int, error) {
	logger.Info("Processando o arquivo do payload")

	tempFile, err := os.CreateTemp("", "uploaded_*.pdf")
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, payload.File)
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao abrir arquivo temporário: %w", err)
	}
	defer file.Close()

	_, err = extractTextFromPDF(file)
	if err != nil {
		return types.CreateAnkiResponse{}, http.StatusInternalServerError, fmt.Errorf("erro ao extrair texto do PDF: %w", err)
	}

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

	// Obter perguntas e respostas usando OpenAI
	answerOpenAi := utils.GetAnswer()

	// Criar estrutura de perguntas a partir da resposta do OpenAI
	questions := parseQuestionsFromOpenAi(answerOpenAi)

	return types.CreateAnkiResponse{
		Question: questions,
	}, http.StatusCreated, nil
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

func parseQuestionsFromOpenAi(answer string) []types.Question {
	var questions []types.Question
	err := json.Unmarshal([]byte(answer), &questions)
	if err != nil {
		logger.LogError("Erro ao parsear resposta do OpenAI:", err)
		return nil
	}
	for i := range questions {
		questions[i].ID = uuid.New()
	}

	return questions
}
