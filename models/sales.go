package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	UserId      uint `gorm:"not null"`
	User        User
	Amount      float64      `json:"price" gorm:"type:decimal(10,2)"`
	SaleDetails []SaleDetail `json:"sale_details"`
}
