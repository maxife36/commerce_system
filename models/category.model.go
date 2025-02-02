package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string    `json:"category" gorm:"not null"`
	Products []Product `json:"products" gorm:"type:varchar(15);constraint:OnUpdate:CASCADE"`
}
