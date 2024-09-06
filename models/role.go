package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role  string `json:"role"`
	Users []User `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
