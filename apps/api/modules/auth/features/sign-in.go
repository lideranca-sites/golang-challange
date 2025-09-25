package features

import (
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type HTTPAuthError struct {
	Error string `json:"error" example:"Mensagem de erro aqui"`
}

const SignInPath = "/sign-in"

type SignInBodyDTO struct {
	Email    *string `validate:"required,email" json:"email"`
	Password *string `validate:"required" json:"password"`
}

// SignIn godoc
// @Summary      Realiza o login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials body SignInBodyDTO true "Credenciais de login"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  HTTPAuthError
// @Router       /auth/sign-in [post]
func SignIn(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Locals("body").(*SignInBodyDTO)
		var user models.User

		result := db.Where("email = ?", *body.Email).First(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(HTTPAuthError{Error: "Invalid credentials"})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*body.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(HTTPAuthError{Error: "Invalid credentials"})
		}

		token, err := CreateJwtToken(CreateJwtTokenDTO{
			UserId: user.ID,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(HTTPAuthError{Error: "Failed to create token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": token,
		})
	}
}
