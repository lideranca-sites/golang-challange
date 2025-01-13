package features

import (
	"crud/modules/database"
	"crud/modules/database/models"

	"github.com/gofiber/fiber/v2"
)

type AddProductDTO struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func AddProduct(c *fiber.Ctx) error {
	var dto AddProductDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user models.User

	if err := database.DB.First(&user, dto.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	product := models.Product{
		UserID: dto.UserID,
		Name:   dto.Name,
	}

	if err := user.AddProduct(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
