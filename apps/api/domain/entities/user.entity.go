package entities

import (
	"errors"
	"example/apps/api/domain/dto"
	"time"
)

type UserEntity struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Products  []ProductEntity
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *UserEntity) AddProduct(product ProductEntity) error {
	if len(u.Products) >= 5 {
		return errors.New("user can't have more than 5 products")
	}

	u.Products = append(u.Products, product)

	return nil
}

func (u *UserEntity) ToDto() dto.UserDTO {
	products := make([]dto.ProductDTO, len(u.Products))
	for i, p := range u.Products {
		products[i] = p.ToDto()
	}

	return dto.UserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Products:  products,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func UserEntityFromDto(i dto.UserDTO) UserEntity {
	products := make([]ProductEntity, len(i.Products))
	for n, p := range i.Products {
		products[n] = ProductEntiryFromDto(p)
	}

	return UserEntity{
		ID:        i.ID,
		Name:      i.Name,
		Email:     i.Email,
		Password:  i.Password,
		Products:  products,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		DeletedAt: i.DeletedAt,
	}
}
