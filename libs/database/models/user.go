package models

import (
	"errors"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt string    `json:"created_at"`
	Products  []Product `json:"products"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt *string   `json:"deleted_at,omitempty"`
}

func (u *User) AddProduct(product Product) error {
	if len(u.Products) >= 5 {
		return errors.New("User can't have more than 5 products")
	}

	u.Products = append(u.Products, product)

	return nil
}
