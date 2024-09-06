package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string            `json:"title" gorm:"not null"`
	Description string            `json:"description"`
	Price       float64           `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int               `json:"stock" gorm:"not null;default:0"`
	CategoryId  *uint             `json:"-"`
	Category    *Category         `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BrandId     *uint             `json:"-"`
	Brand       *Brand            `json:"brand" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Suppliers   []ProductSupplier `json:"suppliers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
