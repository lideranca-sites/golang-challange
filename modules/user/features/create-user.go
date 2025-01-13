package features

import (
	"crud/modules/database"
	"crud/modules/database/models"

	"github.com/gofiber/fiber/v2"
)

type CreateUserDTO struct {
	Name  *string `validate:"required" json:"name"`
	Email *string `validate:"required,email" json:"email"`
}

func CreateUser(c *fiber.Ctx) error {
	var dto CreateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if dto.Name == nil || dto.Email == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and Email are required",
		})
	}

	result := database.DB.Create(&models.User{
		Name:  *dto.Name,
		Email: *dto.Email,
	})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
