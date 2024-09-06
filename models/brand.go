package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Brand string `json:"brand"`
}
