package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func LoadDatabase() *pgx.Conn {

	connStr := os.Getenv(("POSTGRES_URL"))

	if connStr == "" {
		log.Fatal("Connection string not present or invalid in .env")
	}

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to Database")

	return conn

}
