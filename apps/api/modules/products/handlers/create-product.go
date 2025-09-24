package handlers

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

type CreateProductDTO struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,gt=0"`
	Quantity int     `json:"quantity" validate:"required,gte=0"`
}

func CreateProduct(c *fiber.Ctx) error {

	body := c.Locals("body").(*CreateProductDTO)
	userId := c.Locals(locals.UserIdLocal).(int)

	product := models.Product{
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
		UserID:   userId,
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
