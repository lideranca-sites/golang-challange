package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ListProducts(c *fiber.Ctx) error {
	var products []models.Product
	var result *gorm.DB

	user_id := c.Query("user_id")

	if user_id != "" {
		parsed_user_id, err := strconv.Atoi(user_id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user_id",
			})
		}
		result = database.DB.Where("user_id = ?", parsed_user_id).Find(&products)
	} else {
		result = database.DB.Find(&products)
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
