package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const MePath = "/me"

// Me godoc
// @Summary      Obtém dados do usuário logado
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]models.User
// @Security     ApiKeyAuth
// @Router       /auth/me [get]
func Me(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals(locals.UserIdLocal).(int)

		var user models.User

		result := db.Where("id = ?", userId).First(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get user",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user": user,
		})
	}
}
