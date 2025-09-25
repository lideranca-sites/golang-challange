package auth

import (
	"example/apps/api/modules/auth/features"
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func validateSignIn(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.SignInBodyDTO{})
}

func validateSignUp(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.SignUpBodyDTO{})
}

func SetupRoutes(app fiber.Router, db *gorm.DB) {
	group := app.Group("/auth")

	group.Post(features.SignInPath, validateSignIn, features.SignIn(db))

	group.Post(features.SignUpPath, validateSignUp, features.SignUp(db))

	group.Get(features.MePath, middleware.JWTProtected, features.Me(db))
}
