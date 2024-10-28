package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"not null;index"`
	Country     string `json:"country" gorm:"not null"`
	State       string `json:"state" gorm:"not null"`
	City        string `json:"city" gorm:"not null"`
	Postcode    string `json:"postcode" gorm:"not null"`
	AddressLine string `json:"address_line" gorm:"not null"`
	Building    string `json:"building"`
}

type AddressID struct {
	ID uint `json:"id" gorm:"not null"`
}
