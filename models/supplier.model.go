package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Company  string `json:"company" gorm:"type:varchar(50);not null"`
	FiscalId uint16 `json:"cuil_cuit" gorm:"unique"`
	Phone    uint16 `json:"phone"`
	Email    string `json:"email" gorm:"type:varchar(50);unique"`
	Address  string `json:"address" gorm:"type:varchar(50)"`
	Dni      uint16 `json:"dni" gorm:"unique;not null"`
	Billable bool   `json:"billable" gorm:"not null;default:false"`
}
