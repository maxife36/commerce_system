package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price" gorm:"type:decimal(10,2)"`
	Stock       int      `json:"stock"`
	CategoryId  uint     `json:"-"`
	Category    Category `json:"category"`
	BrandId     uint     `json:"-"`
	Brand       Brand    `json:"brand"`
}
