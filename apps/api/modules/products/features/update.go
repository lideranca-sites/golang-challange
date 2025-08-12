package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductBodyDTO struct {
	Name     *string `json:"name" validate:"required_without_all=Price Quantity"`
	Price    *int    `json:"price" validate:"required_without_all=Name Quantity"`
	Quantity *int    `json:"quantity" validate:"required_without_all=Name Price"`
}

func UpdateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*UpdateProductBodyDTO)
	productId := c.Params("id")

	if productId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is required",
		})
	}

	if _, err := strconv.Atoi(productId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	var product models.Product
	if err := database.DB.First(&product, productId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if body.Name != nil {
		product.Name = *body.Name
	}

	if body.Price != nil {
		if *body.Price < 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid product price",
			})
		}
		product.Price = *body.Price
	}

	if body.Quantity != nil {
		if *body.Quantity < 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid product quantity",
			})
		}
		product.Quantity = *body.Quantity
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}
