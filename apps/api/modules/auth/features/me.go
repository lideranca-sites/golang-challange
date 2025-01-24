package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const MePath = "/me"

func Me(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	user_id := c.Locals(locals.UserIdLocal).(int)

	var user models.User

	result := db.Where("id = ?", user_id).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}
