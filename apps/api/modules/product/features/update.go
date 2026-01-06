package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const UpdateProductPath = "/:id"

type UpdateProductDTO struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,gt=0"`
	Quantity int     `json:"quantity" validate:"required,gte=0"`
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is required",
		})
	}

	updateData := new(UpdateProductDTO)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing body",
		})
	}

	product := new(models.Product)
	result := database.DB.First(product, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	result = database.DB.Model(product).Updates(map[string]interface{}{
		"name":     updateData.Name,
		"price":    updateData.Price,
		"quantity": updateData.Quantity,
	})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	database.DB.First(product, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": product,
	})
}
