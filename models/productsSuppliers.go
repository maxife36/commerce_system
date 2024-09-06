package models

import "gorm.io/gorm"

type ProductSupplier struct {
	gorm.Model
	SupplierId uint     `json:"-"`
	Supplier   Supplier `json:"supplier" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductId  uint     `json:"-"`
	Product    Product  `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
