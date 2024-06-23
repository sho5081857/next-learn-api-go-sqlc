package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func NewDB() *pgx.Conn {

	if os.Getenv("GO_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PUBLISHED_PORT"), os.Getenv("DB_DATABASE"))

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(ctx)

	fmt.Println("Connected")

	return conn
}
