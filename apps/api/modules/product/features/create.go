package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const CreatePath = "/"

type CreateProductBodyDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required,gt=0" json:"price"`
	Quantity *int     `validate:"required,gt=0" json:"quantity"`
	UserID   *int     `validate:"required" json:"user_id"`
}

func Create(c *fiber.Ctx) error {
	var body CreateProductBodyDTO
	UserID := c.Locals("user_id").(int)

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	product := models.Product{
		Name:     *body.Name,
		Price:    float64(*body.Price),
		Quantity: *body.Quantity,
		UserID:   int(UserID),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}
