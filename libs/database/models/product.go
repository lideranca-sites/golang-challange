package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       int     `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	UserID   int     `json:"user_id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
