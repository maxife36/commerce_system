package models

type ProductionDetails struct {
	CustomModel
	ProductionId uint       `json:"-" gorm:"primaryKey"`
	Production   Production `json:"production" gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductionId;references:ID"`
	ProductId    uint       `json:"-" gorm:"primaryKey"`
	Product      Product    `json:"product" gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductId;references:ID"`
	Quantity     float64    `json:"quantity" gorm:"not null"`
	Comment      string     `json:"comment" gorm:"type:varchar(100);not null"`
}
