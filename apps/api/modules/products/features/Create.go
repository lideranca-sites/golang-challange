package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const CreateProductPath = "/products"

type ProductBodyDTO struct {
	Name     *string `json:"name" validate:"required"`
	Price    *int    `json:"price" validate:"required"`
	Quantity *int    `json:"quantity" validate:"required"`
}

func CreateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*ProductBodyDTO)
	UserID := c.Locals("user_id").(int)

	product := models.Product{
		Name:     *body.Name,
		Price:    *body.Price,
		Quantity: *body.Quantity,
		UserID:   int(UserID),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
		})

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}
