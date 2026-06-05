package features

import (
	"errors"
	"strconv"

	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const DeleteProductPath = "/:id"

func DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	var product models.Product

	productIdInt, err := strconv.Atoi(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	result := database.DB.First(&product, productIdInt)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server error",
		})

	}

	userId := c.Locals(locals.UserIdLocal).(int)

	if userId != product.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	result = database.DB.Delete(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
