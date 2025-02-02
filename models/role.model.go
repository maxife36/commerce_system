package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role  string `json:"role" gorm:"type:varchar(50);not null"`
	Users []User `json:"users" gorm:"constraint:OnUpdate:CASCADE"`
}
