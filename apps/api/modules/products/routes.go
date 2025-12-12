package products

import (
	"example/apps/api/domain/services"
	gorm_repositories "example/apps/api/infra/database/gorm"
	"example/apps/api/modules/common/middleware"
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	group := app.Group("/products")

	logger := slog.Default()
	userRepo, err := gorm_repositories.NewUserGormRepository()
	if err != nil {
		log.Fatal(err)
	}
	repository, err := gorm_repositories.NewProductGormRepository()
	if err != nil {
		log.Fatal(err)
	}

	service := services.NewProductsServices(repository, userRepo, logger)
	controller := NewProductsController(service)

	group.Get(DEFAULT_PRODUCT_ROUTE, middleware.JWTProtected, controller.ListAll)
	group.Post(DEFAULT_PRODUCT_ROUTE, middleware.JWTProtected, ValidateProductBodyMiddleware, controller.New)
	group.Put(PRODUCT_ROUTE_ID, middleware.JWTProtected, controller.UpdateProduct)
	group.Delete(PRODUCT_ROUTE_ID, middleware.JWTProtected, controller.DeleteProduct)
}
