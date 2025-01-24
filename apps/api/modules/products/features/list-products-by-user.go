package features

import (
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const ListProductByUserPath = "/products"

func ListProductByUser(c *fiber.Ctx) error {
	var result *gorm.DB

	db := c.Locals("db").(*gorm.DB)

	user_id := c.Query("user_id")

	var products []models.Product

	if user_id != "" {
		result = db.Where("user_id = ?", user_id).Find(&products)
	} else {
		result = db.Find(&products)
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to list products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}
