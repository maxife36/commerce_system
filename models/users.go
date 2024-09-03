package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	DNI      uint16 `json:"dni" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Phone    uint16 `json:"phone"`
	Address  string `json:"address"`
	RoleId   uint   `json:"-"`
	Role     Role   `json:"role"`
}
