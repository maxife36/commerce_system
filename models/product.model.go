package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Description string    `json:"description" gorm:"type:varchar(50);not null"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Cost        float64   `json:"cost" gorm:"type:decimal(10,2);not null"`
	Unit        int       `json:"unit" gorm:"constraint:OnUpdate:CASCADE"` /*  ENUMS -> 0: Unidad, 1: Kg, 2: Litro, 3: gr,	4: ml ...etc */
	Stock       float64   `json:"stock" gorm:"not null;default:0"`
	Type        int       `json:"type" gorm:"constraint:OnUpdate:CASCADE"` /* ENUMS -> 0: Product, 1: Combo, 2: Production */
	Condition   int       `json:"condition" gorm:"constraint:OnUpdate:CASCADE"` /* ENUMS -> 0: Compra y venta, 1: Compra, 2: Venta */
	CategoryId  *uint     `json:"-"`
	Category    *Category `json:"category" gorm:"constraint:OnUpdate:CASCADE"`
	BrandId     *uint     `json:"-"`
	Brand       *Brand    `json:"brand" gorm:"constraint:OnUpdate:CASCADE"`
}
