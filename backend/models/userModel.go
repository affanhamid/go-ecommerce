package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string    `json:"email" validate:"required,email" gorm:"not null"`
	Password     string    `json:"password" validate:"required" gorm:"not null"`
	CountryCode  string    `json:"country_code" validate:"required" gorm:"not null"`
	PhoneNumber  string    `json:"phone_number" validate:"required,len=10,numeric" gorm:"not null"`
	FirstName    string    `json:"first_name" validate:"required,min=2,max=100" gorm:"not null"`
	LastName     string    `json:"last_name" validate:"required,min=2,max=100" gorm:"not null"`
	Token        string    `json:"token"`
	UserType     string    `json:"user_type" gorm:"not null" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken string    `json:"refresh_token"`
	Addresses    []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type UserEmail struct {
	Email string `json:"email" validate:"required,email"`
}
