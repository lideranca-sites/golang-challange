package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product id",
		})
	}

	user_id := c.Locals(locals.UserIdLocal).(int)

	var product models.Product
	result := database.DB.Where("id = ? AND user_id = ?", id, user_id).First(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	result = database.DB.Delete(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
