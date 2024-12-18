-- name: CreatePdf :one
INSERT INTO pdfs (user_id, filename, text_content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateStatus :exec
UPDATE pdfs
SET status = $1
WHERE id = $2; 

-- name: UpdateStatusAndText :exec
UPDATE pdfs
SET text_content = $1,
    status =$2
WHERE id = $3;
