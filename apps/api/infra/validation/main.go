package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Error bool
	Field string
	Tag   string
	Value interface{}
}

var validate = validator.New()

func MapErrorMessages(error ValidationError) string {
	var message string

	switch error.Tag {
	case "required":
		message = "The " + error.Field + " field is required"
	case "email":
		message = "The " + error.Field + " field must be a valid email"
	default:
		message = "The " + error.Field + " field is invalid"
	}

	return message
}

func Validate(data interface{}, c *fiber.Ctx) fiber.Map {
	var error ValidationError

	errs := validate.Struct(data)

	if errs != nil {
		first := errs.(validator.ValidationErrors)[0]

		error = ValidationError{
			Error: true,
			Field: first.Field(),
			Tag:   first.Tag(),
			Value: first.Value(),
		}

		return fiber.Map{
			"field":   error.Field,
			"tag":     error.Tag,
			"message": MapErrorMessages(error),
		}
	}

	return nil

}
