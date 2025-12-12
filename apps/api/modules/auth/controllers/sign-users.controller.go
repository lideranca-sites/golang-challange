package auth_controllers

import (
	"example/apps/api/domain/dto"
	"example/apps/api/domain/services"
	"example/apps/api/modules/auth/locals"

	"github.com/gofiber/fiber/v2"
)

const SIGN_UP_PATH = "/sign-up"
const SIGN_IN_PATH = "/sign-in"
const ME_PATH = "/me"

type SignInBodyDTO struct {
	Email    *string `validate:"required,email" json:"email"`
	Password *string `validate:"required" json:"password"`
}

type SignUpBodyDTO struct {
	Name     *string `validate:"required" json:"name"`
	Email    *string `validate:"required,email" json:"email"`
	Password *string `validate:"required" json:"password"`
}

type SignUsersController struct {
	service services.SignUserServices
}

func NewSignUsersController(service services.SignUserServices) SignUsersController {
	return SignUsersController{service: service}
}

func (s SignUsersController) SignIn(c *fiber.Ctx) error {
	var body SignInBodyDTO
	c.BodyParser(&body)

	input := dto.UserDTO{
		Email:    *body.Email,
		Password: *body.Password,
	}
	result, err := s.service.SignIn(input)
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": result,
	})
}

func (s *SignUsersController) SignUp(c *fiber.Ctx) error {
	body := c.Locals("body").(*SignUpBodyDTO)

	input := dto.UserDTO{
		Email:    *body.Email,
		Password: *body.Password,
		Name:     *body.Name,
	}

	result, err := s.service.SignUp(input)
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"access_token": result,
	})
}

func (s *SignUsersController) Me(c *fiber.Ctx) error {
	user_id := c.Locals(locals.UserIdLocal).(int)
	user, err := s.service.FindMe(uint(user_id))
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}
