package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const UpdatePath = "/:id"

type UpdateProductBodyDTO struct {
	Name     *string `validate:"required" json:"name"`
	Price    *int    `validate:"required" json:"price"`
	Quantity *int    `validate:"required" json:"quantity"`
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	body := c.Locals("body").(*UpdateProductBodyDTO)

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if body.Name != nil {
		product.Name = *body.Name
	}
	if body.Price != nil {
		product.Price = *body.Price
	}
	if body.Quantity != nil {
		product.Quantity = *body.Quantity
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to Update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}
