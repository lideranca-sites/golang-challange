package main

import (
	"crud/modules/database"
	"crud/modules/user"

	"github.com/go-playground/validator/v10"
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

	app := fiber.New()

	var validate = validator.New()

	user.SetupRoutes(app)

	app.Listen(":3000")
}
