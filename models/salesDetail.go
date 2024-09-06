package models

import "gorm.io/gorm"

type SaleDetail struct {
	gorm.Model
	SaleId     uint      `json:"-"`
	Title      string    `json:"title" gorm:"not null"`
	Price      float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Quantity   int       `json:"quantity" gorm:"not null"`
	CategoryId *uint     `json:"-"`
	Category   *Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BrandId    *uint     `json:"-"`
	Brand      *Brand    `json:"brand" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
