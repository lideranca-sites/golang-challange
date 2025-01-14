package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SignInPath = "/sign-in"

type SignInBodyDTO struct {
	Email    *string `validate:"required,email" json:"email"`
	Password *string `validate:"required" json:"password"`
}

func SignIn(c *fiber.Ctx) error {
	body := c.Locals("body").(*SignInBodyDTO)

	var user models.User

	result := database.DB.Where("email = ?", *body.Email).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*body.Password))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := CreateJwtToken(CreateJwtTokenDTO{
		UserId: user.ID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
	})

}
