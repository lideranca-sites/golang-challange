package auth

import (
	"example/apps/api/modules/auth/features"
	"example/apps/api/modules/auth/middleware"
	"example/apps/api/validation"

	"github.com/gofiber/fiber/v2"
)

func validateSignIn(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.SignInBodyDTO{})
}

func validateSignUp(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.SignUpBodyDTO{})
}

func SetupRoutes(app fiber.Router) {
	group := app.Group("/auth")

	group.Post(features.SignInPath, validateSignIn, features.SignIn)

	group.Post(features.SignUpPath, validateSignUp, features.SignUp)

	group.Get(features.MePath, middleware.JWTProtected, features.Me)

}
