// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: invoices.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInvoice = `-- name: CreateInvoice :exec
INSERT INTO invoices (customer_id, amount, status, date)
VALUES ($1, $2, $3, $4)
`

type CreateInvoiceParams struct {
	CustomerID pgtype.UUID `json:"customer_id"`
	Amount     int32       `json:"amount"`
	Status     string      `json:"status"`
	Date       pgtype.Date `json:"date"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) error {
	_, err := q.db.Exec(ctx, createInvoice,
		arg.CustomerID,
		arg.Amount,
		arg.Status,
		arg.Date,
	)
	return err
}

const deleteInvoice = `-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1
`

func (q *Queries) DeleteInvoice(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteInvoice, id)
	return err
}

const getFilteredInvoices = `-- name: GetFilteredInvoices :many
SELECT invoices.id,
    invoices.amount,
    invoices.date,
    invoices.status,
    customers.name,
    customers.email,
    customers.image_url
FROM invoices
    JOIN customers ON invoices.customer_id = customers.id
WHERE customers.name ILIKE $1
    OR customers.email ILIKE $1
    OR invoices.amount::text ILIKE $1
    OR invoices.date::text ILIKE $1
    OR invoices.status ILIKE $1
ORDER BY invoices.date DESC
LIMIT $2 OFFSET $3
`

type GetFilteredInvoicesParams struct {
	Name   string `json:"name"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetFilteredInvoicesRow struct {
	ID       pgtype.UUID `json:"id"`
	Amount   int32       `json:"amount"`
	Date     pgtype.Date `json:"date"`
	Status   string      `json:"status"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	ImageUrl string      `json:"image_url"`
}

func (q *Queries) GetFilteredInvoices(ctx context.Context, arg GetFilteredInvoicesParams) ([]GetFilteredInvoicesRow, error) {
	rows, err := q.db.Query(ctx, getFilteredInvoices, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFilteredInvoicesRow
	for rows.Next() {
		var i GetFilteredInvoicesRow
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.Date,
			&i.Status,
			&i.Name,
			&i.Email,
			&i.ImageUrl,
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

const getInvoiceById = `-- name: GetInvoiceById :one
SELECT invoices.id,
    invoices.customer_id,
    invoices.amount,
    invoices.status
FROM invoices
WHERE invoices.id = $1
`

type GetInvoiceByIdRow struct {
	ID         pgtype.UUID `json:"id"`
	CustomerID pgtype.UUID `json:"customer_id"`
	Amount     int32       `json:"amount"`
	Status     string      `json:"status"`
}

func (q *Queries) GetInvoiceById(ctx context.Context, id pgtype.UUID) (GetInvoiceByIdRow, error) {
	row := q.db.QueryRow(ctx, getInvoiceById, id)
	var i GetInvoiceByIdRow
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.Amount,
		&i.Status,
	)
	return i, err
}

const getInvoiceCount = `-- name: GetInvoiceCount :one
SELECT COUNT(*)
FROM invoices
`

func (q *Queries) GetInvoiceCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getInvoiceCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getInvoiceStatusCount = `-- name: GetInvoiceStatusCount :one
SELECT SUM(
        CASE
            WHEN status = 'paid' THEN amount
            ELSE 0
        END
    ) AS "paid",
    SUM(
        CASE
            WHEN status = 'pending' THEN amount
            ELSE 0
        END
    ) AS "pending"
FROM invoices
`

type GetInvoiceStatusCountRow struct {
	Paid    int64 `json:"paid"`
	Pending int64 `json:"pending"`
}

func (q *Queries) GetInvoiceStatusCount(ctx context.Context) (GetInvoiceStatusCountRow, error) {
	row := q.db.QueryRow(ctx, getInvoiceStatusCount)
	var i GetInvoiceStatusCountRow
	err := row.Scan(&i.Paid, &i.Pending)
	return i, err
}

const getInvoicesPages = `-- name: GetInvoicesPages :one
SELECT COUNT(*)
FROM invoices
    JOIN customers ON invoices.customer_id = customers.id
WHERE customers.name ILIKE $1
    OR customers.email ILIKE $1
    OR invoices.amount::text ILIKE $1
    OR invoices.date::text ILIKE $1
    OR invoices.status ILIKE $1
`

func (q *Queries) GetInvoicesPages(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, getInvoicesPages, name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getLatestInvoices = `-- name: GetLatestInvoices :many
SELECT invoices.amount,
    customers.name,
    customers.image_url,
    customers.email,
    invoices.id
FROM invoices
    JOIN customers ON invoices.customer_id = customers.id
ORDER BY invoices.date DESC
LIMIT 5
`

type GetLatestInvoicesRow struct {
	Amount   int32       `json:"amount"`
	Name     string      `json:"name"`
	ImageUrl string      `json:"image_url"`
	Email    string      `json:"email"`
	ID       pgtype.UUID `json:"id"`
}

func (q *Queries) GetLatestInvoices(ctx context.Context) ([]GetLatestInvoicesRow, error) {
	rows, err := q.db.Query(ctx, getLatestInvoices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLatestInvoicesRow
	for rows.Next() {
		var i GetLatestInvoicesRow
		if err := rows.Scan(
			&i.Amount,
			&i.Name,
			&i.ImageUrl,
			&i.Email,
			&i.ID,
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

const updateInvoice = `-- name: UpdateInvoice :exec
UPDATE invoices
SET customer_id = $2,
    amount = $3,
    status = $4
WHERE id = $1
`

type UpdateInvoiceParams struct {
	ID         pgtype.UUID `json:"id"`
	CustomerID pgtype.UUID `json:"customer_id"`
	Amount     int32       `json:"amount"`
	Status     string      `json:"status"`
}

func (q *Queries) UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) error {
	_, err := q.db.Exec(ctx, updateInvoice,
		arg.ID,
		arg.CustomerID,
		arg.Amount,
		arg.Status,
	)
	return err
}
