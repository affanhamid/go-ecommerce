package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID      uint   `json:"user_id" validate:"required" gorm:"not null;index"`
	Country     string `json:"country" validate:"required" gorm:"not null"`
	State       string `json:"state" validate:"required" gorm:"not null"`
	City        string `json:"city" validate:"required" gorm:"not null"`
	Postcode    string `json:"postcode" validate:"required" gorm:"not null"`
	AddressLine string `json:"address_line" validate:"required" gorm:"not null"`
	Building    string `json:"building" validate:"required"`
}

type AddressID struct {
	ID uint `json:"id" gorm:"not null"`
}
