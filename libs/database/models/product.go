package models

import (
	"example/apps/api/domain/dto"
	"example/apps/api/domain/entities"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	UserID   uint
	Price    float64 `gorm:"type:decimal(10,2)"`
	Quantity int
}

func (p Product) ToEntity() entities.ProductEntity {
	return entities.ProductEntity{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
		UserID:   p.UserID,
	}
}

func ProductModelFromDto(input dto.ProductDTO) Product {
	return Product{
		ID:       input.ID,
		Name:     input.Name,
		Price:    input.Price,
		Quantity: input.Quantity,
		UserID:   input.UserID,
	}
}
