package features

import (
	"example/libs/database"
	"example/libs/database/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ListProducts(c *fiber.Ctx) error {
	userIdQuery := c.Query("user_id")
	databaseQuery := database.DB.Model(&models.Product{})

	if userIdQuery != "" {
		userId, err := convertUserIdToInt(userIdQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		databaseQuery = databaseQuery.Where("user_id = ?", userId)
	}

	var products []models.Product
	if err := databaseQuery.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	// remove deleted_at field
	type ProductResponse struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		Quantity  int       `json:"quantity"`
		UserID    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	productsResponse := make([]ProductResponse, 0, len(products))
	for i := range products {
		productsResponse = append(productsResponse, ProductResponse{
			ID:        products[i].ID,
			Name:      products[i].Name,
			Price:     products[i].Price,
			Quantity:  products[i].Quantity,
			UserID:    products[i].UserID,
			CreatedAt: products[i].CreatedAt,
			UpdatedAt: products[i].UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": productsResponse,
	})
}

func convertUserIdToInt(userId string) (int, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}
	return id, nil
}
