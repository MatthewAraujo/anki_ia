// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package repository

import (
	"context"
)

const finAllUsers = `-- name: FinAllUsers :many
select  from user
`

type FinAllUsersRow struct {
}

func (q *Queries) FinAllUsers(ctx context.Context) ([]FinAllUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, finAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FinAllUsersRow
	for rows.Next() {
		var i FinAllUsersRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, name, email, password, created_at 
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, name, email, password, created_at FROM users
WHERE id = $1
`

func (q *Queries) FindUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const insertUsers = `-- name: InsertUsers :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password, created_at
`

type InsertUsersParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) InsertUsers(ctx context.Context, arg InsertUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUsers, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
