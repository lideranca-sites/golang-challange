package features

import (
	"example/apps/api/modules/auth/middleware"
	validation "example/apps/api/modules/auth/products/middleware"

	"github.com/gofiber/fiber/v2"
)

func ValidateBody(c *fiber.Ctx) error {
	return validation.ValidateBodyProduct(c, &ProductBodyDTO{})
}

func Routes(app fiber.Router) {
	group := app.Group("")

	group.Get(GetProductsPath, GetProducts)
	group.Post(CreateProductPath, middleware.JWTProtected, ValidateBody, CreateProduct)
	group.Delete(DeleteProductPath, middleware.JWTProtected, DeleteProduct)
	group.Put(UpdateProductPath, middleware.JWTProtected, UpdateProduct)
}
