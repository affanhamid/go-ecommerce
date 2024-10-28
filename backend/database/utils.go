package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

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
