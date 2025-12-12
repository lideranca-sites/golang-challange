package products

import (
	"example/apps/api/domain/dto"
	"example/apps/api/infra/validation"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ValidateProductBodyMiddleware(c *fiber.Ctx) error {
	fmt.Println("Validatig body?")
	return validation.ValidateBody(c, &dto.ProductDTO{})
}
