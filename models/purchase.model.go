package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	Amount     float64           `json:"amount" gorm:"type:decimal(10,2);not null"`
	Billable   bool              `json:"billable" gorm:"not null;default:false"`
	UserId     uint              `json:"-"`
	User       User              `json:"user" gorm:"constraint:OnUpdate:CASCADE"`
	SupplierId uint              `json:"-"`
	Supplier   Supplier          `json:"supplier" gorm:"constraint:OnUpdate:CASCADE"`
	Details    []PurchaseDetails `json:"details" gorm:"constraint:OnUpdate:CASCADE"`
}
