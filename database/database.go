package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func Connect() *pgx.Conn {

	connStr := os.Getenv(("POSTGRES_URL"))

	if connStr == "" {
		log.Fatal("Connection string not present or invalid in .env")
	}

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn

}

func runSQLFile(file string) string {
	conn := Connect()
	defer conn.Close(context.Background())

	sqlBytes, err := os.ReadFile(fmt.Sprintf("./scripts/%s.sql", file))
	if err != nil {
		log.Fatalf("(%s) Unable to read file: %v", file, err)
	}

	sqlQuery := string(sqlBytes)

	if _, err := conn.Exec(context.Background(), sqlQuery); err != nil {
		log.Fatalf("(%s) Failed to execute SQL query: %v", file, err)
	}

	return fmt.Sprintf("(%s) Successful", file)
}

func CreateTables() string {
	return runSQLFile("create")
}

func DropAllTables() string {
	return runSQLFile("dropAll")
}
