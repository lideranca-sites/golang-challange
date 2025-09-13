package features

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"example/apps/api/modules/products/dtos"
	"example/libs/database"
	"example/libs/database/models"
)

func UpdateProduct(c *fiber.Ctx) error {
	// id na rota
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var body dtos.ProductBodyDTO
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}
	if body.Name == "" || body.Price < 0 || body.Quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validation failed",
		})
	}

	var p models.Product
	if err := database.DB.First(&p, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	p.Name = body.Name
	p.Price = body.Price
	p.Quantity = body.Quantity

	if err := database.DB.Save(&p).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "db error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": dtos.ToDTO(p),
	})
}
