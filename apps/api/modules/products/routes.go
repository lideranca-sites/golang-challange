package products

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/products/features"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	g := app.Group("/products")

	g.Get("/", features.GetProducts)

	g.Post("/", middleware.JWTProtected, features.CreateProduct)
	g.Put("/:id", middleware.JWTProtected, features.UpdateProduct)
	g.Delete("/:id", middleware.JWTProtected, features.DeleteProduct)
}

func Register(r fiber.Router) { SetupRoutes(r) }
