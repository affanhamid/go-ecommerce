package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/affanhamid/go-ecommerce/models"
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

func runSQLFile(DB *gorm.DB, file string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(dir, "../scripts", file+".sql")
	fmt.Println(path)
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("(%s) Unable to read file: %v", file, err)
		return err
	}

	sqlQuery := string(sqlBytes)

	if err := DB.Exec(sqlQuery).Error; err != nil {
		log.Fatalf("(%s) Failed to execute SQL query: %v", file, err)
		return err
	}

	log.Println("Successfully ran SQL file")
	return nil
}

func CreateTables(DB *gorm.DB) error {
	return runSQLFile(DB, "create")
}

func DropTables(DB *gorm.DB) error {
	return runSQLFile(DB, "dropAll")
}

func CreateUser(user *models.User, DB *gorm.DB) error {
	result := DB.Create(&user)
	return result.Error

}

func DeleteUser(user *models.DeleteUser, DB *gorm.DB) error {
	result := DB.Where("email = ?", user.Email).Delete(&models.User{})
	return result.Error
}
