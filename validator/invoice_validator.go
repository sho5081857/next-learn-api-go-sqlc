package validator

import (
	"next-learn-go-sqlc/db/sqlc"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IInvoiceValidator interface {
	InvoiceValidate(invoice sqlc.Invoice) error
}

type invoiceValidator struct{}

func NewInvoiceValidator() IInvoiceValidator {
	return &invoiceValidator{}
}

func (tv *invoiceValidator) InvoiceValidate(invoice sqlc.Invoice) error {
	return validation.ValidateStruct(&invoice,
		validation.Field(
			&invoice.CustomerID,
			validation.Required.Error("CustomerId is required"),
		),
		validation.Field(
			&invoice.Amount,
			validation.Required.Error("Amount is required"),
		),
		validation.Field(
			&invoice.Status,
			validation.Required.Error("Status is required"),
			validation.In(invoice.Status, "pending", "paid").Error("Status must be pending or paid"),
		),
	)
}
