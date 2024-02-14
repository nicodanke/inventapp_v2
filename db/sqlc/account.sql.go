// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: account.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
    code, company_name, email
) VALUES (
    $1, $2, $3
) RETURNING id, code, company_name, phone, email, web_url, active, created_at, updated_at
`

type CreateAccountParams struct {
	Code        string `json:"code"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Code, arg.CompanyName, arg.Email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.CompanyName,
		&i.Phone,
		&i.Email,
		&i.WebUrl,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, code, company_name, phone, email, web_url, active, created_at, updated_at FROM account
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.CompanyName,
		&i.Phone,
		&i.Email,
		&i.WebUrl,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, code, company_name, phone, email, web_url, active, created_at, updated_at FROM account
ORDER BY company_name
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.CompanyName,
			&i.Phone,
			&i.Email,
			&i.WebUrl,
			&i.Active,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account
SET
    company_name = COALESCE($1, company_name),
    phone = COALESCE($2, phone),
    email = COALESCE($3, email),
    web_url = COALESCE($4, web_url),
    active = COALESCE($5, active),
    updated_at = COALESCE($6, updated_at)
WHERE
    id = $7
RETURNING id, code, company_name, phone, email, web_url, active, created_at, updated_at
`

type UpdateAccountParams struct {
	CompanyName pgtype.Text        `json:"company_name"`
	Phone       pgtype.Text        `json:"phone"`
	Email       pgtype.Text        `json:"email"`
	WebUrl      pgtype.Text        `json:"web_url"`
	Active      pgtype.Bool        `json:"active"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	ID          int64              `json:"id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount,
		arg.CompanyName,
		arg.Phone,
		arg.Email,
		arg.WebUrl,
		arg.Active,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.CompanyName,
		&i.Phone,
		&i.Email,
		&i.WebUrl,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}