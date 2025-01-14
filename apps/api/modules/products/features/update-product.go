package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const UpdateProductPath = "/products/:id"

type UpdateProductBodyDTO struct {
	Name     *string `validate:"required" json:"name"`
	Price    *int    `validate:"required" json:"price"`
	Quantity *int    `validate:"required" json:"quantity"`
}

func UpdateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*UpdateProductBodyDTO)
	user_id := c.Locals("user_id").(int)

	found := models.Product{}

	result := database.DB.Where("id = ? AND user_id = ?", c.Params("id"), user_id).First(&found)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	found.Name = *body.Name
	found.Price = *body.Price
	found.Quantity = *body.Quantity

	result = database.DB.Save(&found)


	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": found,
	})
}
