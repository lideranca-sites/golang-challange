package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

type CreateProductBodyDTO struct {
	Name     string  `validate:"required" json:"name"`
	Price    float64 `validate:"required" json:"price"`
	Quantity int     `validate:"required" json:"quantity"`
}

func CreateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*CreateProductBodyDTO)
	userId := c.Locals(locals.UserIdLocal).(int)

	if body.Price < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product price",
		})
	}

	if body.Quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product quantity",
		})
	}

	product := models.Product{
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
		UserID:   userId,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": fiber.Map{
			"id":         product.ID,
			"name":       product.Name,
			"price":      product.Price,
			"quantity":   product.Quantity,
			"user_id":    product.UserID,
			"created_at": product.CreatedAt,
			"updated_at": product.UpdatedAt,
		},
	})
}
