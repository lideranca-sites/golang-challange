package product

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/product/features"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	group := app.Group("/products")

	group.Get("/", features.ListProducts)

	group.Post("/", middleware.JWTProtected, features.CreateProduct)

	group.Put("/:id", middleware.JWTProtected, features.UpdateProduct)

	group.Delete("/:id", middleware.JWTProtected, features.DeleteProduct)
}
