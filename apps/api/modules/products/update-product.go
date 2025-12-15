package products

import (
    "strconv"
    
    "example/libs/database"
    "example/libs/database/models"
    
    "github.com/gofiber/fiber/v2"
)

type UpdateProductBodyDTO struct {
    Name     *string  `validate:"required" json:"name"`
    Price    *float64 `validate:"required" json:"price"`
    Quantity *int     `validate:"required" json:"quantity"`
}

func UpdateProduct(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
    }

    body := c.Locals("body").(*UpdateProductBodyDTO)

    var product models.Product

    result := database.DB.First(&product, id)

    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
    }

    product.Name = *body.Name
    product.Price = *body.Price
    product.Quantity = *body.Quantity

    save := database.DB.Save(&product)

    if save.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Product updated successfully",
        "product": product,
    })
}
