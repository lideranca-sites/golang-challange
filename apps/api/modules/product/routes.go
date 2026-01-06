package product

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/product/features"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	group := app.Group("/product")

	group.Get(features.ListProductPath, features.List)

	group.Post(features.CreateProductPath, middleware.JWTProtected, features.CreateProduct)

	group.Delete(features.DeleteProductPath, middleware.JWTProtected, features.DeleteProduct)

	group.Put(features.UpdateProductPath, middleware.JWTProtected, features.UpdateProduct)
}
