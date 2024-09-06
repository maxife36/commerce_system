package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Dni      uint16 `json:"dni" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Phone    uint16 `json:"phone"`
	Address  string `json:"address"`
	RoleId   uint   `json:"-"`
	Role     Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Sales    []Sale `json:"sales" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
