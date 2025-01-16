package features

import (
	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const DeleteProductPath = "/delete/:id"

func Delete(c *fiber.Ctx) error {
	productID := c.Params("id")
	userID := c.Locals(locals.UserIdLocal).(int)

	// Verifica se o produto existe e pertence ao usuário
	var product models.Product
	if err := database.DB.First(&product, "id = ? AND user_id = ?", productID, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Produto não encontrado ou você não está autorizado a deletar este produto",
		})
	}

	// Deleta o produto
	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao deletar o produto",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Produto " + productID + " deletado com sucesso",
	})
}
