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

func validateCreateProduct(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.AddProductBodyDTO{})
}

func validateEditProduct(c *fiber.Ctx) error {
	return validation.ValidateBody(c, &features.EditProductBodyDTO{})
}

func SetupRoutes(app fiber.Router) {
	SetupAuthRoutes(app)
	SetupProductRoutes(app)
}

func SetupAuthRoutes(app fiber.Router) {
	authGroup := app.Group("/auth")

	authGroup.Post(features.SignInPath, validateSignIn, features.SignIn)
	authGroup.Post(features.SignUpPath, validateSignUp, features.SignUp)
	authGroup.Get(features.MePath, middleware.JWTProtected, features.Me)
}

func SetupProductRoutes(app fiber.Router) {
	productGroup := app.Group("/product")
	productGroup.Get(features.ListProductsPath, middleware.JWTProtected, features.ListProducts)
	productGroup.Post(features.CreateProductPath, middleware.JWTProtected, validateCreateProduct, features.Add)
	productGroup.Put(features.EditProductPath, middleware.JWTProtected, validateEditProduct, features.Edit)
	productGroup.Delete(features.DeleteProductPath, middleware.JWTProtected, features.Delete)
	productGroup.Get(features.GetProductPath, middleware.JWTProtected, features.GetProduct)

}
