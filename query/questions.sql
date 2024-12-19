-- name: CreateQuestion :one
INSERT INTO questions (pdf_id, question_text)
VALUES ($1, $2)
RETURNING id;

-- name: GetQuestionsByPdfId :many
SELECT id, pdf_id, question_text
FROM questions
WHERE pdf_id = $1;
