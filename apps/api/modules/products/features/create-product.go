package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const CreateProductPath = ""

type CreateProductDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required" json:"price"`
	Quantity *int     `validate:"required" json:"quantity"`
}

func CreateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*CreateProductDTO)

	user_id := c.Locals("user_id").(int)

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
		"message": "Product created successfully",
		"product": product,
	})
}
