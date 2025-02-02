// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateInvoice(ctx context.Context, arg CreateInvoiceParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteInvoice(ctx context.Context, id pgtype.UUID) error
	GetAllCustomers(ctx context.Context) ([]GetAllCustomersRow, error)
	GetAllRevenues(ctx context.Context) ([]Revenue, error)
	GetCustomerCount(ctx context.Context) (int64, error)
	GetFilteredCustomers(ctx context.Context, name string) ([]GetFilteredCustomersRow, error)
	GetFilteredInvoices(ctx context.Context, arg GetFilteredInvoicesParams) ([]GetFilteredInvoicesRow, error)
	GetInvoiceById(ctx context.Context, id pgtype.UUID) (GetInvoiceByIdRow, error)
	GetInvoiceCount(ctx context.Context) (int64, error)
	GetInvoiceStatusCount(ctx context.Context) (GetInvoiceStatusCountRow, error)
	GetInvoicesPages(ctx context.Context, name string) (int64, error)
	GetLatestInvoices(ctx context.Context) ([]GetLatestInvoicesRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id pgtype.UUID) (User, error)
	UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) error
}

var _ Querier = (*Queries)(nil)
