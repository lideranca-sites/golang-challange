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

	group.Get(features.GetProductsPath, features.GetProducts)
	group.Post(features.CreateProductPath, middleware.JWTProtected, validateCreateProduct, features.CreateProduct)
	group.Put(features.UpdateProductPath, middleware.JWTProtected, validateUpdateProduct, features.UpdateProduct)
	group.Delete(features.DeleteProductPath, middleware.JWTProtected, features.DeleteProduct)
}
