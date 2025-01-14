package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBody(c *fiber.Ctx, body interface{}) error {
	var error ValidationError

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errs := validate.Struct(body)

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

	c.Locals("body", body)

	return c.Next()
}

// 🚨 nao usar essa funcao, o go fiber ta com problema 🚨
// Dica: pode usar c.ParamsParser(&params) direto no da feature, pesquise como usar
// 
// func ValidateParams(c *fiber.Ctx, params interface{}) error {
// 	var error ValidationError

// 	if err := c.ParamsParser(&params); err != nil {
// 		fmt.Println(params)
// 		fmt.Println(reflect.TypeOf(params))
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	errs := validate.Struct(params)

// 	if errs != nil {
// 		first := errs.(validator.ValidationErrors)[0]

// 		error = ValidationError{
// 			Error: true,
// 			Field: first.Field(),
// 			Tag:   first.Tag(),
// 			Value: first.Value(),
// 		}

// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"field":   error.Field,
// 			"tag":     error.Tag,
// 			"message": MapErrorMessages(error),
// 		})
// 	}

// 	c.Locals("params", params)

// 	return c.Next()
// }

// 🚨 nao usar essa funcao, o go fiber ta com problema 🚨
// Dica: pode usar c.QueryParser(&query) direto no da feature, pesquise como usar
// func ValidateQuery(c *fiber.Ctx, query interface{}) error {
// 	var error ValidationError

// 	if err := c.QueryParser(&query); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	errs := validate.Struct(query)

// 	if errs != nil {
// 		first := errs.(validator.ValidationErrors)[0]

// 		error = ValidationError{
// 			Error: true,
// 			Field: first.Field(),
// 			Tag:   first.Tag(),
// 			Value: first.Value(),
// 		}

// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"field":   error.Field,
// 			"tag":     error.Tag,
// 			"message": MapErrorMessages(error),
// 		})
// 	}

// 	c.Locals("query", query)

// 	return c.Next()
// }
