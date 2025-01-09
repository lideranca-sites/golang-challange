package features

import (
	"crud/modules/database"
	"crud/modules/database/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	database.DB.Delete(&models.User{}, id)

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
