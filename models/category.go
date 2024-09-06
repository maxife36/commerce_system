package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string    `json:"category"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
