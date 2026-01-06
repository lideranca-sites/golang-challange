package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	userID := c.Query("user_id")

	var products []models.Product

	if userID != "" {
		result := database.DB.Where("user_id = ?", userID).Find(&products)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch products",
			})
		}
	} else {
		result := database.DB.Find(&products)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch products",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}