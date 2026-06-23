package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const GetPath = "/"

func Get(c *fiber.Ctx) error {
	var products []models.Product

	userID := c.Query("user_id")
	query := database.DB

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get products",
		})
	}

	return c.JSON(fiber.Map{
		"products": products,
	})
}
