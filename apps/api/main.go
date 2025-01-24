package main

import (
	"example/apps/api/infra/server"
	"example/libs/database"
	"example/libs/sqs"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()
	sqs := sqs.CreateClient()

	app := server.Setup(db, sqs)

	app.Listen(":3000")
}
