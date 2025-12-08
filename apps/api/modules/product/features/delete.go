package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const DeletePath = "/:id"

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to Delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
