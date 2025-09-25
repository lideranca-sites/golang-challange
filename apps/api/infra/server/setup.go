package server

import (
	"example/apps/api/modules/auth"
	"example/apps/api/modules/products"
	"example/apps/api/modules/products/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swag "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, productHandler *handlers.ProductHandler) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/swagger/*", swag.WrapHandler)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth.SetupRoutes(v1, db)
	products.SetupRoutes(v1, productHandler)

	return app
}
