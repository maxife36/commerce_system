package models

import (
	"gorm.io/gorm"
	"time"
)

type CustomModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
