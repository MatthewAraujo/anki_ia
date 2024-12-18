-- name: CreatePdf :one
INSERT INTO pdfs (user_id, filename, status, text_content)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: UpdateStatus :one
UPDATE pdfs
SET status = $1
WHERE id = $2 
RETURNING id;
