package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	GetProductPath   = "/:id"
	ListProductsPath = "/list"
)

func GetProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	userID := c.Locals(locals.UserIdLocal).(int)

	var product models.Product
	if err := database.DB.First(&product, "id = ? AND user_id = ?", productID, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Produto não encontrado ou você não está autorizado a visualizar este produto",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":         product.ID,
		"name":       product.Name,
		"user_id":    product.UserID,
		"price":      product.Price,
		"quantity":   product.Quantity,
		"created_at": product.CreatedAt.Format(time.RFC3339),
		"updated_at": product.UpdatedAt.Format(time.RFC3339),
	})
}

func ListProducts(c *fiber.Ctx) error {
	userID := c.Locals(locals.UserIdLocal).(int)

	var products []models.Product
	if err := database.DB.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao listar os produtos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": func() []fiber.Map {
			mappedProducts := make([]fiber.Map, len(products))
			for i, product := range products {
				mappedProducts[i] = fiber.Map{
					"id":         product.ID,
					"name":       product.Name,
					"user_id":    product.UserID,
					"price":      product.Price,
					"quantity":   product.Quantity,
					"created_at": product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
					"updated_at": product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
				}
			}
			return mappedProducts
		}(),
	})
}
