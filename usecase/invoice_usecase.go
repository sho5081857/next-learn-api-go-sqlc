package usecase

import (
	"context"
	"next-learn-go-sqlc/db/sqlc"
	"next-learn-go-sqlc/validator"

	"github.com/jackc/pgx/v5/pgtype"
)

type IInvoiceUsecase interface {
	GetLatestInvoices() ([]sqlc.GetLatestInvoicesRow, error)
	GetFilteredInvoices(query string, offset, limit int32) ([]sqlc.GetFilteredInvoicesRow, error)
	GetInvoiceCount() (int64, error)
	GetInvoiceStatusCount() (int64, int64, error)
	GetInvoicesPages(query string) (int64, error)
	GetInvoiceById(invoiceId pgtype.UUID) (sqlc.GetInvoiceByIdRow, error)
	CreateInvoice(invoice sqlc.Invoice) error
	UpdateInvoice(invoice sqlc.Invoice) error
	DeleteInvoice(invoiceId pgtype.UUID) error
}

type invoiceUsecase struct {
	iq sqlc.Querier
	iv validator.IInvoiceValidator
}

func NewInvoiceUsecase(iq sqlc.Querier, iv validator.IInvoiceValidator) IInvoiceUsecase {
	return &invoiceUsecase{iq, iv}
}

func (iu *invoiceUsecase) GetLatestInvoices() ([]sqlc.GetLatestInvoicesRow, error) {
	invoices, err := iu.iq.GetLatestInvoices(context.Background())
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (iu *invoiceUsecase) GetFilteredInvoices(query string, offset, limit int32) ([]sqlc.GetFilteredInvoicesRow, error) {
	arg := sqlc.GetFilteredInvoicesParams{Name: query, Offset: offset, Limit: limit}
	invoices, err := iu.iq.GetFilteredInvoices(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (iu *invoiceUsecase) GetInvoiceCount() (int64, error) {
	count, err := iu.iq.GetInvoiceCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iu *invoiceUsecase) GetInvoiceStatusCount() (int64, int64, error) {
	i, err := iu.iq.GetInvoiceStatusCount(context.Background())
	if err != nil {
		return 0, 0, err
	}
	return i.Pending, i.Paid, nil
}

func (iu *invoiceUsecase) GetInvoicesPages(query string) (int64, error) {
	count, err := iu.iq.GetInvoicesPages(context.Background(), query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iu *invoiceUsecase) GetInvoiceById(invoiceId pgtype.UUID) (sqlc.GetInvoiceByIdRow, error) {
	invoice, err := iu.iq.GetInvoiceById(context.Background(), invoiceId)
	if err != nil {
		return sqlc.GetInvoiceByIdRow{}, err
	}

	return invoice, nil
}

func (iu *invoiceUsecase) CreateInvoice(invoice sqlc.Invoice) error {
	if err := iu.iv.InvoiceValidate(invoice); err != nil {
		return err
	}
	arg := sqlc.CreateInvoiceParams{Amount: invoice.Amount, CustomerID: invoice.CustomerID, Status: invoice.Status}

	if err := iu.iq.CreateInvoice(context.Background(), arg); err != nil {
		return err
	}

	return nil
}

func (iu *invoiceUsecase) UpdateInvoice(invoice sqlc.Invoice) error {
	if err := iu.iv.InvoiceValidate(invoice); err != nil {
		return err
	}

	arg := sqlc.UpdateInvoiceParams{Amount: invoice.Amount, CustomerID: invoice.CustomerID, Status: invoice.Status}
	if err := iu.iq.UpdateInvoice(context.Background(), arg); err != nil {
		return err
	}
	return nil
}

func (iu *invoiceUsecase) DeleteInvoice(invoiceId pgtype.UUID) error {
	if err := iu.iq.DeleteInvoice(context.Background(), invoiceId); err != nil {
		return err
	}
	return nil
}
