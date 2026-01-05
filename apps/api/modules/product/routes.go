package product

import (
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/modules/product/features"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	group := app.Group("/product")

	group.Get("/",  middleware.JWTProtected, features.FindByUser)

}
