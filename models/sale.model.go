package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	Amount   float64       `json:"price" gorm:"type:decimal(10,2);not null"`
	Billable bool          `json:"billable" gorm:"not null;default:false"`
	UserId   uint          `json:"-"`
	User     User          `json:"user" gorm:"constraint:OnUpdate:CASCADE"`
	ClientId uint          `json:"-"`
	Client   Client        `json:"client" gorm:"constraint:OnUpdate:CASCADE"`
	Details  []SaleDetails `json:"Details" gorm:"constraint:OnUpdate:CASCADE"`
}
