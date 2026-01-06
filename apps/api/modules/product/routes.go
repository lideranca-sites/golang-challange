package product

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/product/features"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	group := app.Group("/product")

	group.Get("/", features.List)

	group.Post("/", middleware.JWTProtected, features.CreateProduct)

	group.Delete("/", middleware.JWTProtected, features.DeleteProduct)

	group.Put("/", middleware.JWTProtected, features.UpdateProduct)
}
