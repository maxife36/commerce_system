package models

import "gorm.io/gorm"

type ProductSupplier struct {
	gorm.Model
	SupplierId uint     `json:"-"`
	Supplier   Supplier `json:"supplier"`
	ProductId  uint     `json:"-"`
	Product    Product  `json:"product"`
}
