package features

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"example/apps/api/modules/products/dtos"
	"example/libs/database"
	"example/libs/database/models"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if uidStr := c.Query("user_id"); uidStr != "" {
		uid, err := strconv.Atoi(uidStr)
		if err != nil || uid <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user_id",
			})
		}
		if err := database.DB.Where("user_id = ?", uid).Find(&products).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "db error",
			})
		}
	} else {
		if err := database.DB.Find(&products).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "db error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": dtos.ToDTOs(products),
	})
}
