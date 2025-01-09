package calculator

import (
	"github.com/gofiber/fiber/v2"
)

type CalculateDTO struct {
	Operation    string `json:"operation"`
	FirstNumber  int    `json:"first_number"`
	SecondNumber int    `json:"second_number"`
}

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/calculate", func(c *fiber.Ctx) error {
		var dto CalculateDTO

		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		calculator := GetCalculator(dto.Operation)

		result := calculator.MakeCalculation(dto.FirstNumber, dto.SecondNumber)

		return c.JSON(fiber.Map{
			"result": result,
		})
	})
}
