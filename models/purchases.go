package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	Details    json.RawMessage `json:"details" gorm:"type:json"`
	Amount     float64         `json:"amount" gorm:"type:decimal(10,2)"`
	UserId     uint            `json:"-"`
	User       User            `json:"user"`
	SupplierId uint            `json:"-"`
	Supplier   Supplier        `json:"supplier"`
}
