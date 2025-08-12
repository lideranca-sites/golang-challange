package models

import (
	"time"

	"gorm.io/gorm"
)

/**
 * @todo: esqueci de add a tag json aos campos do gorm, por exemplo: CreatedAt.
 * Entao esses campos estao como camelCase enquanto o restante esta em snake_case.
 */
type Product struct {
	gorm.Model
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Price     float64    `gorm:"type:double" json:"price"`
	Quantity  int        `json:"quantity"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
