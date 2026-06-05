package dto

import "time"

type UserDTO struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"-"`
	CreatedAt time.Time    `json:"created_at"`
	Products  []ProductDTO `json:"products"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt time.Time    `json:"deleted_at,omitempty"`
}
