package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ValidateBody(c *fiber.Ctx, body interface{}) error {
	var error ValidationError

	log.Printf("ValidateBody: Starting validation for path: %s, method: %s", c.Path(), c.Method())
	log.Printf("ValidateBody: Content-Type: %s", c.Get("Content-Type"))
	log.Printf("ValidateBody: Content-Length: %s", c.Get("Content-Length"))

	// Read body first to see if it's there
	bodyBytes := c.Body()
	log.Printf("ValidateBody: Body length: %d bytes", len(bodyBytes))
	if len(bodyBytes) > 0 {
		log.Printf("ValidateBody: Body content: %s", string(bodyBytes))
	}

	if err := c.BodyParser(&body); err != nil {
		log.Printf("ValidateBody: BodyParser error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("ValidateBody: Body parsed successfully")

	errs := validate.Struct(body)

	if errs != nil {
		log.Printf("ValidateBody: Validation errors found: %v", errs)
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

	log.Printf("ValidateBody: Validation passed, setting body in locals")
	c.Locals("body", body)

	log.Printf("ValidateBody: Calling next handler")
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
