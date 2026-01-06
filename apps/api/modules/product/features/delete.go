package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const DeleteProductPath = "/"

func DeleteProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	id := c.Query("id")

	result := database.DB.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"deleted_at": product.DeletedAt,
	})
}
