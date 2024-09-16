package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role  string `json:"role" gorm:"not null"`
	Users []User `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
