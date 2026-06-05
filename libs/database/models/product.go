package models

import (
	"time"
)

type Product struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Quantity  int        `json:"quantity"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
