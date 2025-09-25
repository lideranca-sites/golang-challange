package models

import "time"

type Product struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Quantity  int        `json:"quantity"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
