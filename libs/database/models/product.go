package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	UserID   int    `json:"user_id"`
}
