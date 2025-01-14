package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const ListProductByUserPath = "/products"

func ListProductByUser(c *fiber.Ctx) error {
	var result *gorm.DB
	
	user_id := c.Query("user_id")
	
	var products []models.Product

	if user_id != "" {
		result = database.DB.Where("user_id = ?", user_id).Find(&products)
	} else {
		result = database.DB.Find(&products)
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
