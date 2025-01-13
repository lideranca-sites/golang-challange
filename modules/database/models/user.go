package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	Products  []Product `json:"products"`
}

func (u *User) AddProduct(product Product) error {
	if len(u.Products) >= 5 {
		return errors.New("User can't have more than 5 products")
	}

	u.Products = append(u.Products, product)

	return nil
}
