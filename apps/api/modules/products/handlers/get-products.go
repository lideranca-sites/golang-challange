package handlers

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	query := database.DB

	userId := c.Query("user_id")
	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}
