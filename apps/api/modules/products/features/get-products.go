package features

import (
	"strconv"

	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
)

const GetProductsPath = "/"

func GetProducts(c *fiber.Ctx) error {

	var products []models.Product
	userId := c.Query("user_id")
	query := database.DB

	if userId != "" {
		userIdInt, err := strconv.Atoi(userId)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user_id",
			})
		}

		query = query.Where("user_id = ?", userIdInt)

	}

	result := query.Find(&products)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
	})

}
