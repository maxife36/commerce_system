package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Brand    string    `json:"brand" gorm:"not null"`
	Products []Product `json:"products" gorm:"type:varchar(15);constraint:OnUpdate:CASCADE"`
}
