package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Brand    string    `json:"brand" gorm:"not null"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
