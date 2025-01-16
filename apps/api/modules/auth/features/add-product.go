package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

const CreateProductPath = "/create"

type AddProductBodyDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required" json:"price"`
	Quantity *int     `validate:"required" json:"quantity"`
}

func Add(c *fiber.Ctx) error {
	body := c.Locals("body").(*AddProductBodyDTO)
	fmt.Println(*body)
	user_id := c.Locals(locals.UserIdLocal).(int)
	fmt.Println(user_id)

	product := models.Product{
		Name:     *body.Name,
		UserID:   user_id,
		Price:    *body.Price,
		Quantity: *body.Quantity,
	}

	result := database.DB.Create(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":         product.ID,
		"name":       product.Name,
		"user_id":    product.UserID,
		"price":      product.Price,
		"quantity":   product.Quantity,
		"created_at": product.CreatedAt.Format(time.RFC3339),
		"updated_at": product.UpdatedAt.Format(time.RFC3339),
	})
}
