package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Brand    string    `json:"brand"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
