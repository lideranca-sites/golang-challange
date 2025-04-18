package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const DeleteProductPath = "/products/:id"

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_id is required",
		})
	}

	if err := database.DB.Delete(&models.Product{}, productID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
