package main

import (
	"next-learn-go-sqlc/infrastructure/database"
	"next-learn-go-sqlc/infrastructure/database/sqlc"
	"next-learn-go-sqlc/router"
	"os"
)

func main() {

	conn := database.NewDB()
	q := sqlc.New(conn)

	e := router.NewRouter(q)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
