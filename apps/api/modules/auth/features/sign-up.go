package features

import (
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SignUpPath = "/sign-up"

type SignUpBodyDTO struct {
	Name     *string `validate:"required" json:"name"`
	Email    *string `validate:"required,email" json:"email"`
	Password *string `validate:"required" json:"password"`
}

// SignUp godoc
// @Summary      Registra um novo usuário
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body SignUpBodyDTO true "Dados de registro"
// @Success      201  {object}  map[string]string
// @Router       /auth/sign-up [post]
func SignUp(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Locals("body").(*SignUpBodyDTO)

		hash, err := bcrypt.GenerateFromPassword([]byte(*body.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to process request",
			})
		}

		user := models.User{
			Name:     *body.Name,
			Email:    *body.Email,
			Password: string(hash),
		}

		result := db.Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		token, err := CreateJwtToken(CreateJwtTokenDTO{
			UserId: user.ID,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create token",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"access_token": token,
		})
	}
}
