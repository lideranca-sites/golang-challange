package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserID   int    `json:"user_id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
