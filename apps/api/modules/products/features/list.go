package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)


const ListProductsPath = "/"

func ListProducts(c *fiber.Ctx) error {
	var products []models.Product
	var filters struct {
		UserId string `query:"user_id"`
	}

	if err := c.QueryParser(&filters); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing query",
		})
	}

	query := database.DB.Model(&models.Product{})

	if filters.UserId != "" {
		query = query.Where("user_id = ?", filters.UserId)
	}

	if result := query.Find(&products); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}




