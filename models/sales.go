package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	User   User
	Amount float32
}
