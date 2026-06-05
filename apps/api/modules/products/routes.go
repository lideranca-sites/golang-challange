package products

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/products/features"
	"example/apps/api/validation"
	"github.com/gofiber/fiber/v2"
)

func validateCreateProduct(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.CreateProductBodyDTO{})
}

func validateUpdateProduct(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.UpdateProductBodyDTO{})
}

func SetupRoutes(app fiber.Router) {
	group := app.Group("/products")

	group.Get("/", features.ListProducts)

	group.Post("/", middleware.JWTProtected, validateCreateProduct, features.CreateProduct)

	group.Put("/:id", middleware.JWTProtected, validateUpdateProduct, features.UpdateProduct)

	group.Delete("/:id", middleware.JWTProtected, features.DeleteProduct)
}
