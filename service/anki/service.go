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

	"github.com/ledongthuc/pdf"
)

type Service struct {
	db *repository.Queries

	dbTx *sql.DB
}

func NewService(db *repository.Queries, dbTx *sql.DB) *Service {
	return &Service{
		db:   db,
		dbTx: dbTx,
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

func (s *Service) CreateAnki(payload *types.CreateAnkiPayload) (string, int, error) {
	logger.Info("Processando o arquivo do payload")

	tempFile, err := os.CreateTemp("", "uploaded_*.pdf")
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, payload.File)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao abrir arquivo temporário: %w", err)
	}
	defer file.Close()

	text, err := extractTextFromPDF(file)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao extrair texto do PDF: %w", err)
	}

	return text, http.StatusCreated, nil
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
