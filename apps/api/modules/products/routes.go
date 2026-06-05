package products

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/products/features"
	"example/apps/api/validation"

	"github.com/gofiber/fiber/v2"
)

func validateCreate(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.CreateProductDTO{})
}

func validateUpdate(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.UpdateProductDTO{})
}

func SetupRoutes(app fiber.Router) {
	group := app.Group("/products")

	group.Post(features.CreateProductPath, middleware.JWTProtected, validateCreate, features.CreateProduct)

	group.Get(features.ListProductsPath, features.ListProducts)

	group.Put(features.UpdateProductPath, middleware.JWTProtected, validateUpdate, features.UpdateProduct)

	group.Delete(features.DeleteProductPath, middleware.JWTProtected, features.DeleteProduct)
}
