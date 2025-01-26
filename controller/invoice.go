package controller

import (
	"net/http"
	"next-learn-go-sqlc/infrastructure/database/sqlc"
	"next-learn-go-sqlc/usecase"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type InvoiceController interface {
	GetLatestInvoices(c echo.Context) error
	GetFilteredInvoices(c echo.Context) error
	GetInvoiceCount(c echo.Context) error
	GetInvoiceStatusCount(c echo.Context) error
	GetInvoicesPages(c echo.Context) error
	GetInvoiceById(c echo.Context) error
	CreateInvoice(c echo.Context) error
	UpdateInvoice(c echo.Context) error
	DeleteInvoice(c echo.Context) error
}

type invoiceController struct {
	iu usecase.InvoiceUseCase
}

func NewInvoiceController(iu usecase.InvoiceUseCase) InvoiceController {
	return &invoiceController{iu}
}

func (ic *invoiceController) GetLatestInvoices(c echo.Context) error {

	invoiceRes, err := ic.iu.GetLatestInvoices()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetFilteredInvoices(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	query := c.QueryParams().Get("query")

	invoiceRes, err := ic.iu.GetFilteredInvoices(query, int32(offset), int32(limit))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetInvoiceCount(c echo.Context) error {
	invoiceRes, err := ic.iu.GetInvoiceCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetInvoiceStatusCount(c echo.Context) error {
	pending, paid, err := ic.iu.GetInvoiceStatusCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]int64{"pending": pending, "paid": paid})
}

func (ic *invoiceController) GetInvoicesPages(c echo.Context) error {
	query := c.QueryParams().Get("query")

	invoiceRes, err := ic.iu.GetInvoicesPages(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetInvoiceById(c echo.Context) error {
	var invoiceId pgtype.UUID
	err := invoiceId.Scan(c.Param("invoiceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	invoiceRes, err := ic.iu.GetInvoiceById(invoiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) CreateInvoice(c echo.Context) error {

	invoice := sqlc.Invoice{}
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := ic.iu.CreateInvoice(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Invoice created")
}

func (ic *invoiceController) UpdateInvoice(c echo.Context) error {
	invoice := sqlc.Invoice{}
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := ic.iu.UpdateInvoice(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Invoice updated")
}

func (ic *invoiceController) DeleteInvoice(c echo.Context) error {
	var invoiceId pgtype.UUID
	err := invoiceId.Scan(c.Param("invoiceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ic.iu.DeleteInvoice(invoiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
