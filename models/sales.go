package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	UserId      uint         `json:"-" gorm:"not null"`
	User        User         `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount      float64      `json:"price" gorm:"type:decimal(10,2)"`
	SaleDetails []SaleDetail `json:"sale_details" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
