package features

import (
	"crud/modules/database"
	"crud/modules/database/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UpdateUser(c *fiber.Ctx) error {
	var dto UpdateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Params("id")

	database.DB.Model(&models.User{}).Where("id = ?", id).Updates(&models.User{
		Name:  dto.Name,
		Email: dto.Email,
	})

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
	})
}
