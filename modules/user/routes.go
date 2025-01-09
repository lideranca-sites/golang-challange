package user

import (
	"crud/modules/user/features"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/users", features.ListUser)
	app.Post("/api/v1/users", features.CreateUser)
	app.Patch("/api/v1/users/:id", features.UpdateUser)
	app.Delete("/api/v1/users/:id", features.DeleteUser)
}
