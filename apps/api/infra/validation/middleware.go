package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBody(c *fiber.Ctx, dto interface{}) error {
	var error ValidationError

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errs := validate.Struct(dto)

	if errs != nil {
		first := errs.(validator.ValidationErrors)[0]

		error = ValidationError{
			Error: true,
			Field: first.Field(),
			Tag:   first.Tag(),
			Value: first.Value(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"field":   error.Field,
			"tag":     error.Tag,
			"message": MapErrorMessages(error),
		})
	}

	c.Locals("body", dto)

	return c.Next()
}

func ValidateParams(c *fiber.Ctx, dto interface{}) error {
	var error ValidationError

	if err := c.ParamsParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errs := validate.Struct(dto)

	if errs != nil {
		first := errs.(validator.ValidationErrors)[0]

		error = ValidationError{
			Error: true,
			Field: first.Field(),
			Tag:   first.Tag(),
			Value: first.Value(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"field":   error.Field,
			"tag":     error.Tag,
			"message": MapErrorMessages(error),
		})
	}

	c.Locals("params", dto)

	return c.Next()
}

func ValidateQuery(c *fiber.Ctx, dto interface{}) error {
	var error ValidationError

	if err := c.QueryParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errs := validate.Struct(dto)

	if errs != nil {
		first := errs.(validator.ValidationErrors)[0]

		error = ValidationError{
			Error: true,
			Field: first.Field(),
			Tag:   first.Tag(),
			Value: first.Value(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"field":   error.Field,
			"tag":     error.Tag,
			"message": MapErrorMessages(error),
		})
	}

	c.Locals("query", dto)

	return c.Next()
}
