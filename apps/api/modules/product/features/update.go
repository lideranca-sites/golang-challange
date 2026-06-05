package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	userID := c.Locals(locals.UserIdLocal).(int)
	if product.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	var updateData models.Product
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Update fields
	product.Name = updateData.Name
	product.Price = updateData.Price
	product.Quantity = updateData.Quantity

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}
