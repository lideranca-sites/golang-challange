package features

import "github.com/gofiber/fiber/v2"

type FindByUserDTO struct {
	Name     *string `validate:"required" json:"name"`
}

func FindByUser(c *fiber.Ctx) error{
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": "teste",
	})
}