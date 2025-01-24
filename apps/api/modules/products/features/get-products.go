package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const GetProductsPath = ""

func GetProducts(c *fiber.Ctx) error {
	user_id := c.Query("user_id")

	products := []models.Product{}

	if user_id != "" {
		database.DB.Where("user_id = ?", user_id).Find(&products)
	} else {
		database.DB.Find(&products)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})
}
