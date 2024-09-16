package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Holder   string            `json:"holder" gorm:"not null"`
	Phone    uint16            `json:"phone"`
	Email    string            `json:"email" gorm:"unique"`
	Address  string            `json:"address"`
	Dni      uint16            `json:"dni" gorm:"unique;not null"`
	FiscalId uint16            `json:"cuil_cuit" gorm:"unique;not null"`
	Products []ProductSupplier `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
