package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	Details    json.RawMessage `json:"details" gorm:"type:json;not null"`
	Amount     float64         `json:"amount" gorm:"type:decimal(10,2);not null"`
	UserId     uint            `json:"-"`
	User       User            `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SupplierId uint            `json:"-"`
	Supplier   Supplier        `json:"supplier" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
