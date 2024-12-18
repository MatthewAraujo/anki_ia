-- name: FinAllUsers :many
select * from user;

-- name: InsertUsers :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * 
FROM users
WHERE email = $1
LIMIT 1;