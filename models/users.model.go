package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(50)"`
	LastName string `json:"last_name" gorm:"type:varchar(50)"`
	Username string `json:"username" gorm:"type:varchar(50);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(60);not null"`
	Dni      uint16 `json:"dni" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"type:varchar(50);unique"`
	Phone    uint16 `json:"phone"`
	Address  string `json:"address" gorm:"type:varchar(50)"`
	RoleId   uint   `json:"-"`
	Role     Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE"`
	Sales    []Sale `json:"sales" gorm:"constraint:OnUpdate:CASCADE"`
}
