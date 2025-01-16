package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

const EditProductPath = "/edit/:id"

type EditProductBodyDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required" json:"price"`
	Quantity *int     `validate:"required" json:"quantity"`
}

func Edit(c *fiber.Ctx) error {
	body := c.Locals("body").(*EditProductBodyDTO)
	productID := c.Params("id")
	userID := c.Locals(locals.UserIdLocal).(int)

	var product models.Product
	if err := database.DB.First(&product, "id = ? AND user_id = ?", productID, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Produto não encontrado ou você não está autorizado a editar este produto",
		})
	}

	// Atualiza os campos do produto com os valores fornecidos
	product.Name = *body.Name
	product.Price = *body.Price
	product.Quantity = *body.Quantity

	// Salva as alterações no banco de dados
	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao atualizar o produto",
		})
	}

	// Retorna o produto atualizado
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
