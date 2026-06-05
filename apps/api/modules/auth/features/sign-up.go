package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"log"

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

	log.Printf("SignUp request for email: %s", *body.Email)

	hash, err := bcrypt.GenerateFromPassword([]byte(*body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user := models.User{
		Name:     *body.Name,
		Email:    *body.Email,
		Password: string(hash),
	}

	log.Println("Attempting to create user in database...")
	result := database.DB.Create(&user)

	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create user",
			"details": result.Error.Error(),
		})
	}

	log.Printf("User created successfully with ID: %d", user.ID)

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
