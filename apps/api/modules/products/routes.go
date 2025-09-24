package products

import (
	"example/apps/api/infra/validation"
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/products/handlers"

	"github.com/gofiber/fiber/v2"
)

func validateCreateProduct(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &handlers.CreateProductDTO{})
}

func validateUpdateProduct(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &handlers.UpdateProductDTO{})
}

func SetupRoutes(router fiber.Router) {
	products := router.Group("/products")

	products.Post("/", middleware.JWTProtected, validateCreateProduct, handlers.CreateProduct)
	products.Get("/", handlers.GetProducts)
	products.Put("/:id", middleware.JWTProtected, validateUpdateProduct, handlers.UpdateProduct)
	products.Delete("/:id", middleware.JWTProtected, handlers.DeleteProduct)
}
