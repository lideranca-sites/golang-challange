package dto

type ProductDTO struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name" validate:"required"`
	UserID   uint    `json:"user_id" validate:"required,gt=0"`
	Price    float64 `json:"price" validate:"required,gt=0"`
	Quantity int     `json:"quantity" validate:"required,gt=0"`
}
