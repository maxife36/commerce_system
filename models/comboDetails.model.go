package models

type ComboDetails struct {
	CustomModel
	ComboId   uint     `json:"-" gorm:"primaryKey"` 
	ProductId uint     `json:"-" gorm:"primaryKey"` 
	Combo     Combo    `json:"combo" gorm:"constraint:OnDelete:CASCADE;foreignKey:ComboId;references:ID"`
	Product   Product  `json:"product" gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductId;references:ID"`
	Quantity  float64  `json:"quantity" gorm:"not null"`
}