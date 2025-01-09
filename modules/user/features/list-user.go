package features

import (
	"crud/modules/database"
	"crud/modules/database/models"
	"github.com/gofiber/fiber/v2"
)

func ListUser(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(fiber.Map{
		"users": users,
	})
}
