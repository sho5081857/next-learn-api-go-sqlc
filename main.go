package main

import (
	"next-learn-go-sqlc/controller"
	"next-learn-go-sqlc/db"
	"next-learn-go-sqlc/db/sqlc"
	"next-learn-go-sqlc/router"
	"next-learn-go-sqlc/usecase"
	"next-learn-go-sqlc/validator"
	"os"
)

func main() {

	conn := db.NewDB()
	q := sqlc.New(conn)

	userValidator := validator.NewUserValidator()
	invoiceValidator := validator.NewInvoiceValidator()

	// userRepository := repository.NewUserRepository(db)
	// invoiceRepository := repository.NewInvoiceRepository(db)

	userUsecase := usecase.NewUserUsecase(q, userValidator)
	invoiceUsecase := usecase.NewInvoiceUsecase(q, invoiceValidator)
	revenueUsecase := usecase.NewRevenueUsecase(q)
	customerUsecase := usecase.NewCustomerUsecase(q)

	userController := controller.NewUserController(userUsecase)
	invoiceController := controller.NewInvoiceController(invoiceUsecase)
	revenueController := controller.NewRevenueController(revenueUsecase)
	customerController := controller.NewCustomerController(customerUsecase)

	e := router.NewRouter(userController, invoiceController, revenueController, customerController)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
