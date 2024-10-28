package database

import (
	"github.com/affanhamid/go-ecommerce/models"
	"gorm.io/gorm"
)

func AddAddress(address *models.Address, DB *gorm.DB) error {
	result := DB.Create(&address)
	return result.Error
}

func GetAddressesFromUser(userId uint, DB *gorm.DB) ([]models.Address, error) {
	var addresses []models.Address
	if err := DB.Where("user_id = ?", userId).Find(&addresses).Error; err != nil {
		return addresses, err
	}
	return addresses, nil
}

func DeleteAddress(addressID *models.AddressID, DB *gorm.DB) error {
	var address models.Address
	if err := DB.Where("id = ?", addressID.ID).Delete(&address).Error; err != nil {
		return err
	}
	DB.Unscoped().Where("id = ?", addressID.ID).Delete(&address)
	return nil
}
