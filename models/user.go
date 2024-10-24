package models

type User struct {
	Email       string `gorm:"primaryKey" validate:"required,email"`
	CountryCode string `json:"country_code" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,len=10,numeric"`
	FirstName   string `json:"first_name" validate:"required,min=2,max=100"`
	LastName    string `json:"last_name" validate:"required,min=2,max=100"`
}

type DeleteUser struct {
	Email string `json:"email" validate:"required,email"`
}
