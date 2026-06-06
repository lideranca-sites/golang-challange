package features

import (
	"strconv"

	"example/apps/api/modules/auth/locals"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const UpdateProductPath = "/:id"

type UpdateProductBodyDTO struct {
	Name     *string  `validate:"required" json:"name"`
	Price    *float64 `validate:"required" json:"price"`
	Quantity *int     `validate:"required" json:"quantity"`
}

func UpdateProduct(c *fiber.Ctx) error {
	body := c.Locals("body").(*UpdateProductBodyDTO)
	productId := c.Params("id")
	var product models.Product

	productIdInt, err := strconv.Atoi(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	product, err = LoadProduct(c, productIdInt)
	if err != nil {
		return err
	}

	userId := c.Locals(locals.UserIdLocal).(int)

	if userId != product.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	product.Name = *body.Name
	product.Price = *body.Price
	product.Quantity = *body.Quantity

	result := database.DB.Save(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})

}
