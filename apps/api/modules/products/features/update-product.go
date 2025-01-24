package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const UpdateProductPath = "/:id"

type UpdateProductDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required" json:"price"`
	Quantity *int     `validate:"required" json:"quantity"`
}

func UpdateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*UpdateProductDTO)

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	product := models.Product{
		ID:       id,
		Name:     *body.Name,
		Price:    *body.Price,
		Quantity: *body.Quantity,
	}

	result := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(&product)

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
