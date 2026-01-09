package main

import (
	"example/apps/api/infra/server"
	"example/libs/database"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Try to load .env from project root (two levels up from apps/api)
	envPath := filepath.Join("..", "..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		if err := godotenv.Load(".env"); err != nil {
			panic("Failed to load .env file. Make sure it exists in the project root (golang-challange/.env)")
		}
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	app := server.Setup()

	log.Println("Starting server on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
