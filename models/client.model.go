package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Company  string `json:"company" gorm:"type:varchar(50);not null"`
	FiscalId uint16 `json:"cuil_cuit"`
	Phone    uint16 `json:"phone"`
	Email    string `json:"email" gorm:"type:varchar(50);unique"`
	Address  string `json:"address" gorm:"type:varchar(50)"`
	Billable   bool              `json:"billable" gorm:"not null;default:false"`
}