package validation

import "github.com/gofiber/fiber/v2"

func ValidateBodyProduct(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Locals("body", body)

	return c.Next()
}
