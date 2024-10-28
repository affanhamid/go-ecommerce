package database

import (
	"github.com/affanhamid/go-ecommerce/models"
	"github.com/affanhamid/go-ecommerce/utils"
	"gorm.io/gorm"
)

func CreateUser(user *models.User, DB *gorm.DB) error {
	result := DB.Create(&user)
	return result.Error

}

func GetUser(userEmail string, DB *gorm.DB) (models.UserInterface, error) {
	var existingUser models.User
	if err := DB.Where("email = ?", userEmail).First(&existingUser).Error; err != nil {
		return utils.ToUserInterface(&existingUser), err
	}

	return utils.ToUserInterface(&existingUser), nil
}

func DeleteUser(user *models.UserEmail, DB *gorm.DB) error {
	var existingUser models.User
	if err := DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return err
	}

	DB.Where("email = ?", user.Email).Delete(&existingUser)
	return nil
}

func DeleteUserPerma(user *models.UserEmail, DB *gorm.DB) error {
	return DeleteUser(user, DB.Unscoped())
}
