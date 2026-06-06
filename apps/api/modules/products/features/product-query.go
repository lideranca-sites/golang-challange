package features

import (
	"errors"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoadProduct(c *fiber.Ctx, id int) (models.Product, error) {
	var product models.Product

	result := database.DB.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return product, c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}

		return product, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server error",
		})
	}

	return product, result.Error

}
