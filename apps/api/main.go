package main

import (
	"example/apps/api/infra/server"
	"example/apps/api/modules/products/handlers"
	"example/apps/api/modules/products/repositories"
	"example/apps/api/modules/products/services"
	_ "example/docs"
	"example/libs/database"
	"log"

	"github.com/joho/godotenv"
)

// @title API de Desafio Go
// @version 1.0
// @description API para gerenciamento de produtos.
// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	productRepository := repositories.NewProductGormRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	app := server.Setup(db, productHandler)

	log.Println("Server is running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
