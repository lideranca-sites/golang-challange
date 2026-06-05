package products

import (
	"example/apps/api/domain/dto"
	"example/apps/api/infra/validation"

	"github.com/gofiber/fiber/v2"
)

func ValidateProductBodyMiddleware(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &dto.ProductDTO{})
}
