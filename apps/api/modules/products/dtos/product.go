package dtos

import "example/libs/database/models"

type ProductBodyDTO struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ProductDTO struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	UserID   int     `json:"user_id"`
}

func ToDTO(m models.Product) ProductDTO {
	return ProductDTO{
		ID:       m.ID,
		Name:     m.Name,
		Price:    m.Price,
		Quantity: m.Quantity,
		UserID:   m.UserID,
	}
}

func ToDTOs(ms []models.Product) []ProductDTO {
	out := make([]ProductDTO, 0, len(ms))
	for _, m := range ms {
		out = append(out, ToDTO(m))
	}
	return out
}
