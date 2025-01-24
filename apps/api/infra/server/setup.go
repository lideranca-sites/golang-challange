package server

import (
	"example/apps/api/modules/auth"
	"example/apps/api/modules/products"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, sqs *sqs.Client) *fiber.App {
	app := fiber.New()
	v1 := app.Group("/api/v1")

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("sqs", sqs)
		c.Locals("db", db)
		return c.Next()
	})

	auth.SetupRoutes(v1)
	products.SetupRoutes(v1)

	return app
}
