package product

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/product/features"
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

	group.Get(features.GetPath, features.Get)
	group.Post(features.CreatePath, middleware.JWTProtected, validateCreateProduct, features.Create)
	group.Put(features.UpdatePath, middleware.JWTProtected, validateUpdateProduct, features.Update)
	group.Delete(features.DeletePath, middleware.JWTProtected, features.Delete)
}
