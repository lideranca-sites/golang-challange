package products

import (
    "example/apps/api/modules/auth/locals"
    "example/libs/database"
    "example/libs/database/models"
    "github.com/gofiber/fiber/v2"
)

type CreateProductBodyDTO struct {
    Name     *string  `validate:"required" json:"name"`
    Price    *float64 `validate:"required" json:"price"`
    Quantity *int     `validate:"required" json:"quantity"`
}

func CreateProduct(c *fiber.Ctx) error {
    body := c.Locals("body").(*CreateProductBodyDTO)

    user_id := c.Locals(locals.UserIdLocal).(int)

    product := models.Product{
        Name:     *body.Name,
        Price:    *body.Price,
        Quantity: *body.Quantity,
        UserID:   user_id,
    }

    result := database.DB.Create(&product)

    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create product",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Product created successfully",
        "product": product,
    })
}
