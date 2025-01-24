package features

import (
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const DeleteProductPath = "/products/:id"

func DeleteProduct(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	user_id := c.Locals("user_id").(int)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	result := db.Where("id = ? AND user_id = ?", id, user_id).Delete(&models.Product{})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
