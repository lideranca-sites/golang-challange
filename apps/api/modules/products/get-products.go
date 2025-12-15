package products

import (
    "example/libs/database"
    "example/libs/database/models"
    
    "github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
    var query struct {
        UserId *int `query:"user_id"`
    }

    if err := c.QueryParser(&query); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    var products []models.Product

    if query.UserId != nil {
        database.DB.Where("user_id = ?", *query.UserId).Find(&products)
    } else {
        database.DB.Find(&products)
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "products": products,
    })
}
