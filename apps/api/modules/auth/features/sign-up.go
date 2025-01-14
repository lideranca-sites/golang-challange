package features

import (
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SignUpPath = "/sign-up"

type SignUpBodyDTO struct {
	Name     *string `validate:"required" json:"name"`
	Email    *string `validate:"required,email" json:"email"`
	Password *string `validate:"required" json:"password"`
}

func SignUp(c *fiber.Ctx) error {
	body := c.Locals("body").(*SignUpBodyDTO)

	hash, err := bcrypt.GenerateFromPassword([]byte(*body.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     *body.Name,
		Email:    *body.Email,
		Password: string(hash),
	}

	result := database.DB.Create(&user)

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
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"access_token": token,
	})
}
