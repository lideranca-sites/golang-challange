package server

import (
	auth "example/apps/api/modules"
	"example/apps/api/modules/products/features"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	features.Routes(v1)
	auth.SetupRoutes(v1)

	return app
}
