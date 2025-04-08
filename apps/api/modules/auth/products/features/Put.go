package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const UpdateProductPath = "/products/:id"

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product models.Product

	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is required",
		})
	}

	if err := database.DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	type UpdateProductDTO struct {
		Name     *string  `json:"name"`
		Price    *float64 `json:"price"`
		Quantity *float64 `json:"quantity"`
	}

	var updateData UpdateProductDTO
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if updateData.Name != nil {
		product.Name = *updateData.Name
	}
	if updateData.Price != nil {
		product.Price = int(*updateData.Price)
	}
	if updateData.Quantity != nil {
		product.Quantity = int(*updateData.Quantity)
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": map[string]interface{}{
			"id":       product.ID,
			"name":     product.Name,
			"price":    product.Price,
			"quantity": product.Quantity,
		},
	})
}
