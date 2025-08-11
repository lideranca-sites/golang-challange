package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListProducts(c *fiber.Ctx) error {
	userIdQuery := c.Query("user_id")
	databaseQuery := database.DB.Model(&models.Product{})

	if userIdQuery != "" {
		userId, err := convertUserIdToInt(userIdQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		databaseQuery = databaseQuery.Where("user_id = ?", userId)
	}

	var products []models.Product
	if err := databaseQuery.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}

func convertUserIdToInt(userId string) (int, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}
	return id, nil
}
