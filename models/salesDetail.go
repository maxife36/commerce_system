package models

import "gorm.io/gorm"

type SaleDetail struct {
	gorm.Model
	SaleId     uint     `json:"-"`
	Title      string   `json:"title"`
	Price      float64  `json:"price" gorm:"type:decimal(10,2)"`
	Quantity   int      `json:"quantity"`
	CategoryId uint     `json:"-"`
	Category   Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BrandId    uint     `json:"-"`
	Brand      Brand    `json:"brand" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
