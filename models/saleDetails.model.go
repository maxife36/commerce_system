package models

import "gorm.io/gorm"

type SaleDetails struct {
	gorm.Model
	Quantity  float64  `json:"quantity" gorm:"not null"`
	SaleId    uint    `json:"-"`
	Sale      Sale    `json:"sale" gorm:"constraint:OnUpdate:CASCADE"`
	ProductId uint    `json:"-"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE"`
}