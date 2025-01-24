package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const DeleteProductPath = "/:id"

func DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	result := database.DB.Delete(&models.Product{}, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
