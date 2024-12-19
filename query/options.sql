-- name: InsertOption :exec
INSERT INTO options (question_id, option_key, option_text, is_correct)
VALUES ($1, $2, $3, $4);

-- name: GetOptionsByQuestionId :many
SELECT id, question_id, option_key, option_text, is_correct
FROM options
WHERE question_id = $1;
