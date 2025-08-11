package products

import (
	"example/apps/api/modules/products/features"

	"github.com/gofiber/fiber/v2"
)

const BASE_PATH = "/products"

func SetupRoutes(router fiber.Router) {
	router.Get(BASE_PATH, features.ListProducts)
}
