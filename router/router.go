package router

import (
	"net/http"
	"next-learn-go-sqlc/controller"
	"next-learn-go-sqlc/controller/middleware"
	"next-learn-go-sqlc/infrastructure/database/sqlc"
	"next-learn-go-sqlc/usecase"
	"next-learn-go-sqlc/validator"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	q *sqlc.Queries,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CorsMiddleware())
	jwtMiddleware := middleware.JwtMiddleware()

	userValidator := validator.NewUserValidator()
	invoiceValidator := validator.NewInvoiceValidator()

	userUseCase := usecase.NewUserUseCase(q, userValidator)
	invoiceUseCase := usecase.NewInvoiceUseCase(q, invoiceValidator)
	revenueUseCase := usecase.NewRevenueUseCase(q)
	customerUseCase := usecase.NewCustomerUseCase(q)

	userController := controller.NewUserController(userUseCase)
	invoiceController := controller.NewInvoiceController(invoiceUseCase)
	revenueController := controller.NewRevenueController(revenueUseCase)
	customerController := controller.NewCustomerController(customerUseCase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.POST("/signup", userController.SignUp)
	e.POST("/login", userController.LogIn)

	i := e.Group("/invoices")
	i.Use(jwtMiddleware)
	i.GET("/latest", invoiceController.GetLatestInvoices)
	i.GET("/filtered", invoiceController.GetFilteredInvoices)
	i.GET("/count", invoiceController.GetInvoiceCount)
	i.GET("/status/count", invoiceController.GetInvoiceStatusCount)
	i.GET("/pages", invoiceController.GetInvoicesPages)
	i.GET("/:invoiceId", invoiceController.GetInvoiceById)
	i.POST("", invoiceController.CreateInvoice)
	i.PATCH("/:invoiceId", invoiceController.UpdateInvoice)
	i.DELETE("/:invoiceId", invoiceController.DeleteInvoice)

	r := e.Group("/revenues")
	r.Use(jwtMiddleware)
	r.GET("", revenueController.GetAllRevenues)

	c := e.Group("/customers")
	c.Use(jwtMiddleware)
	c.GET("", customerController.GetAllCustomers)
	c.GET("/filtered", customerController.GetFilteredCustomers)
	c.GET("/count", customerController.GetCustomerCount)

	u := e.Group("/user")
	u.Use(jwtMiddleware)
	u.GET("", userController.GetUserById)
	u.GET("/email", userController.GetUserByEmail)
	return e
}
