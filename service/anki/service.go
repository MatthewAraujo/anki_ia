package anki

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/types"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
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
	defer tempFile.Close()

	_, err = io.Copy(tempFile, payload.File)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	file, err := os.Open(tempFile.Name())
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao abrir conteúdo do arquivo: %w", err)
	}

	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("erro ao abrir conteúdo do arquivo: %w", err)
	}

	texExtractor, err := extractor.New(pdfReader)

	return texExtractor, http.StatusCreated, nil
}
