package products

import (
    "example/apps/api/modules/auth/middleware"
    "example/apps/api/validation"
    
    "github.com/gofiber/fiber/v2"
)

func validateCreate(c *fiber.Ctx) error {
    return validation.ValidateBody(c, &CreateProductBodyDTO{})
}

func validateUpdate(c *fiber.Ctx) error {
    return validation.ValidateBody(c, &UpdateProductBodyDTO{})
}

func SetupRoutes(app fiber.Router) {
    group := app.Group("/products")

    group.Get("", GetProducts)

    group.Post("", middleware.JWTProtected, validateCreate, CreateProduct)

    group.Put(":id", middleware.JWTProtected, validateUpdate, UpdateProduct)

    group.Delete(":id", middleware.JWTProtected, DeleteProduct)
}
