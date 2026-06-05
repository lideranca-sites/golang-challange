package products

import (
    "strconv"
    
    "example/libs/database"
    "example/libs/database/models"

    "github.com/gofiber/fiber/v2"
)

func DeleteProduct(c *fiber.Ctx) error {
    idStr := c.Params("id")
    id, err := strconv.Atoi(idStr)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
    }

    result := database.DB.Delete(&models.Product{}, id)

    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
