-- name: InsertOption :exec
INSERT INTO options (question_id, option_key, option_text, is_correct)
VALUES ($1, $2, $3, $4);
