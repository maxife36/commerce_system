package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Holder   string `json:"holder"`
	Phone    uint16 `json:"phone"`
	Email    string `json:"email" gorm:"unique"`
	Address  string `json:"address"`
	Dni      uint16 `json:"dni" gorm:"unique"`
	FiscalId uint16 `json:"cuil_cuit" gorm:"unique"`
}
