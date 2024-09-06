package models

import "gorm.io/gorm"

type SaleDetail struct {
	gorm.Model
	SaleId   uint    `json:"-"`
	Title    string  `json:"title"`
	Price    float64 `json:"price" gorm:"type:decimal(10,2)"`
	Quantity int     `json:"quantity"`
	//CategoryId uint
	//Category Category
	//BrandId uint
	//Brand Brand
}
