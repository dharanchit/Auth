package models

import "gorm.io/gorm"

type UserGo struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
