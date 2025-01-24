package main

import (
	"example/apps/api/infra/server"
	"example/libs/database"
)

func main() {
	if err := database.Connect(); err != nil {
		panic(err)
	}

	app := server.Setup()

	app.Listen(":3000")
}
