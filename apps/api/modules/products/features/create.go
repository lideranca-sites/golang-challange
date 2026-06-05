package features

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

const CreateProductPath = "/"

func CreateProduct(c *fiber.Ctx) error {
	userId := c.Locals(locals.UserIdLocal).(int)
	p := new(models.Product)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing body",
		})
	}


	p.UserID = userId
	result := database.DB.Create(&p)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"product": p,
	})
}

