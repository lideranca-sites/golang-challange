package auth

import (
	"example/apps/api/domain/services"
	gorm_repositories "example/apps/api/infra/database/gorm"
	auth_controllers "example/apps/api/modules/auth/controllers"
	features "example/apps/api/modules/auth/controllers"
	"example/apps/api/modules/common/middleware"
	"example/apps/api/utils"
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

	repository, err := gorm_repositories.NewUserGormRepository()
	if err != nil {
		panic(err)
	}

	crypto := utils.NewCryptoUtils()
	token := utils.NewJwtUtils()
	service := services.NewSignUserServices(repository, &crypto, &token)
	controller := auth_controllers.NewSignUsersController(service)

	group.Post(auth_controllers.SIGN_IN_PATH, validateSignIn, controller.SignIn)

	group.Post(auth_controllers.SIGN_UP_PATH, validateSignUp, controller.SignUp)

	group.Get(auth_controllers.ME_PATH, middleware.JWTProtected, controller.Me)

}
