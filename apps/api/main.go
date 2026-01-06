package main

import (
	"example/apps/api/infra/server"
	"example/libs/database"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	app := server.Setup()

	fmt.Println("Server is running on port 3000")

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
