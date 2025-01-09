package features

import (
	"crud/modules/database"
	"crud/modules/database/models"

	"github.com/gofiber/fiber/v2"
)

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(c *fiber.Ctx) error {
	var dto CreateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Create(&models.User{
		Name:  dto.Name,
		Email: dto.Email,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
