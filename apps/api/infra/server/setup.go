package server

import (
	"example/apps/api/modules/auth"
	"example/apps/api/modules/products"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Printf("=== INCOMING REQUEST ===")
		log.Printf("Method: %s", c.Method())
		log.Printf("Path: %s", c.Path())
		log.Printf("OriginalURL: %s", c.OriginalURL())
		log.Printf("IP: %s", c.IP())
		log.Printf("========================")
		err := c.Next()
		log.Printf("=== REQUEST COMPLETED ===")
		return err
	})

	// Add a simple test route to verify server is working
	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("✅ Test route / hit!")
		return c.JSON(fiber.Map{
			"message": "Server is running!",
			"status":  "ok",
		})
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		log.Println("✅ Test route /test hit!")
		return c.JSON(fiber.Map{
			"message": "Test endpoint works!",
			"path":    c.Path(),
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth.SetupRoutes(v1)
	products.SetupRoutes(v1)

	log.Println("Routes registered:")
	log.Println("  GET  /")
	log.Println("  POST /api/v1/auth/sign-up")
	log.Println("  POST /api/v1/auth/sign-in")
	log.Println("  GET  /api/v1/auth/me")
	log.Println("  GET  /api/v1/products")
	log.Println("  POST /api/v1/products")
	log.Println("  PUT  /api/v1/products/:id")
	log.Println("  DELETE /api/v1/products/:id")

	return app
}
