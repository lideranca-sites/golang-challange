package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const GetProductsPath = "/products"

func GetProducts(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	var products []models.Product

	result := database.DB

	if userID == "" {
		result = result.Find(&products)
	} else {
		result = result.Where("user_id = ?", userID).Find(&products)
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}
