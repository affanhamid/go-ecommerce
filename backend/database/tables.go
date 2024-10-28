package database

import (
	"github.com/affanhamid/go-ecommerce/models"
	"gorm.io/gorm"
)

func CreateTables(DB *gorm.DB) error {
	if err := DB.AutoMigrate(&models.User{}, &models.Address{}); err != nil {
		return err
	}

	return nil
}

func DropTables(DB *gorm.DB) error {
	return runSQLFile(DB, "dropAll")
}
