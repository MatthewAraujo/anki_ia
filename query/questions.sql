-- name: CreateQuestion :one
INSERT INTO questions (pdf_id, question_text)
VALUES ($1, $2)
RETURNING id;

