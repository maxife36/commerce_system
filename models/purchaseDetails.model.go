package models

import "gorm.io/gorm"

type PurchaseDetails struct {
	gorm.Model
	Quantity   float64  `json:"quantity" gorm:"not null"`
	PurchaseId uint     `json:"-"`
	Purchase   Purchase `json:"purchase" gorm:"constraint:OnUpdate:CASCADE"`
	ProductId  uint     `json:"-"`
	Product    Product  `json:"product" gorm:"constraint:OnUpdate:CASCADE"`
}
