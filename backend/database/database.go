package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not open .env, trying file path ../.env " + err.Error())
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("Could not find .env file")
		}
	}

	connStr := os.Getenv("POSTGRES_URL")

	if connStr == "" {
		log.Fatal("Connection string not present or invalid in .env")
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return db

}
