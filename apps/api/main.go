package main

import (
	"example/apps/api/infra/server"
	"example/libs/database"

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

	app.Listen(":3000")
}
