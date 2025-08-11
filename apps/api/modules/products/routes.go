package products

import (
	"example/apps/api/infra/validation"
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/products/features"

	"github.com/gofiber/fiber/v2"
)

func validateCreateProductRequest(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.CreateProductBodyDTO{})
}

const BASE_PATH = "/products"

func SetupRoutes(router fiber.Router) {
	router.Get(BASE_PATH, features.ListProducts)
	router.Post(BASE_PATH, middleware.JWTProtected, validateCreateProductRequest, features.CreateProduct)
}
