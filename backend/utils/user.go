package utils

import "github.com/affanhamid/go-ecommerce/models"

func ToUserInterface(user *models.User) models.UserInterface {
	return models.UserInterface{
		ID:          user.ID,
		Email:       user.Email,
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
	}
}
