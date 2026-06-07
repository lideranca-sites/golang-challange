package features

type ProductBodyDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required" json:"price"`
	Quantity *int     `validate:"required" json:"quantity"`
}