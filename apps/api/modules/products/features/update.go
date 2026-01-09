package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type UpdateProductBodyDTO struct {
	Name     *string  `json:"name"`
	Price    *float64 `json:"price"`
	Quantity *int     `json:"quantity"`
}

func UpdateProduct(c *fiber.Ctx) error {
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

	body := c.Locals("body").(*UpdateProductBodyDTO)

	if body.Name != nil {
		product.Name = *body.Name
	}
	if body.Price != nil {
		product.Price = *body.Price
	}
	if body.Quantity != nil {
		product.Quantity = *body.Quantity
	}

	result = database.DB.Save(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}
