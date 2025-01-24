package server

import (
	"example/apps/api/modules/auth"
	"example/apps/api/modules/products"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth.SetupRoutes(v1)
	products.SetupRoutes(v1)

	return app
}
