// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customers.sql

package repository

import (
	"context"
)

const finAllCustomers = `-- name: FinAllCustomers :many
select id, name, email, password, role from customers
`

func (q *Queries) FinAllCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, finAllCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Role,
		); err != nil {
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

const findCustomerByEmail = `-- name: FindCustomerByEmail :one
SELECT id, name, email, password, role 
FROM customers
WHERE email = $1
LIMIT 1
`

func (q *Queries) FindCustomerByEmail(ctx context.Context, email string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, findCustomerByEmail, email)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const findCustomerByID = `-- name: FindCustomerByID :one
SELECT id, name, email, password, role FROM customers
WHERE id = $1
`

func (q *Queries) FindCustomerByID(ctx context.Context, id int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, findCustomerByID, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const insertCustomers = `-- name: InsertCustomers :one
INSERT INTO customers (name, email, password, role)
VALUES ($1, $2, $3, $4)
RETURNING id, name, email, password, role
`

type InsertCustomersParams struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

func (q *Queries) InsertCustomers(ctx context.Context, arg InsertCustomersParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, insertCustomers,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Role,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Role,
	)
	return i, err
}