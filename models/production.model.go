package models

type Production struct {
	CustomModel
	ID       uint    `json:"-" gorm:"primaryKey"`
	Product  Product `json:"product" gorm:"constraint:OnDelete:CASCADE;foreignKey:ID;references:ID"`
	Quantity float64 `json:"quantity" gorm:"not null"`
	Comment  string  `json:"comment" gorm:"type:varchar(100);not null"`
}
