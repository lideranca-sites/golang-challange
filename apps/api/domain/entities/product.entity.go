package entities

import (
	"example/apps/api/domain/dto"
	"time"
)

type ProductEntity struct {
	ID        uint
	Name      string
	UserID    uint
	Price     float64
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (p *ProductEntity) ToDto() dto.ProductDTO {
	return dto.ProductDTO{
		ID:       p.ID,
		Name:     p.Name,
		UserID:   p.UserID,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
}

func ProductEntiryFromDto(i dto.ProductDTO) ProductEntity {
	return ProductEntity{
		ID:       i.ID,
		Name:     i.Name,
		UserID:   i.UserID,
		Price:    i.Price,
		Quantity: i.Quantity,
	}
}
