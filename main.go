package main

import (
	"crud/modules/database"
	"crud/modules/user"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			return c.Status(code).SendString("something happened at the disco!")
		},
	})

	user.SetupRoutes(app)

	app.Listen(":3000")
}
