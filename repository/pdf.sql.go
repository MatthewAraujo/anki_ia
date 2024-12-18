// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: pdf.sql

package repository

import (
	"context"
	"database/sql"
)

const createPdf = `-- name: CreatePdf :one
INSERT INTO pdfs (user_id, filename, status, text_content)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreatePdfParams struct {
	UserID      int32          `json:"user_id"`
	Filename    string         `json:"filename"`
	Status      sql.NullString `json:"status"`
	TextContent sql.NullString `json:"text_content"`
}

func (q *Queries) CreatePdf(ctx context.Context, arg CreatePdfParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createPdf,
		arg.UserID,
		arg.Filename,
		arg.Status,
		arg.TextContent,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateStatus = `-- name: UpdateStatus :one
UPDATE pdfs
SET status = $1
WHERE id = $2 
RETURNING id
`

type UpdateStatusParams struct {
	Status sql.NullString `json:"status"`
	ID     int32          `json:"id"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateStatus, arg.Status, arg.ID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
