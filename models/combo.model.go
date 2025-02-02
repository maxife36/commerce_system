package models

import "time"

type Combo struct {
	CustomModel
	ID       uint      `json:"id" gorm:"primaryKey"`
	Product  Product   `json:"product" gorm:"constraint:OnDelete:CASCADE;foreignKey:ID;references:ID"`
	DateFrom time.Time `json:"date_from" gorm:"type:date;not null"`
	DateTo   time.Time `json:"date_to" gorm:"type:date;not null"`
}
