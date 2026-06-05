package middleware

import (
	"example/apps/api/modules/auth/locals"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		TokenLookup: "header:Authorization",
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey:  locals.JwtLocal,
		AuthScheme:  "Bearer",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals(locals.JwtLocal).(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			id := claims["id"].(float64)

			c.Locals(locals.UserIdLocal, int(id))

			return c.Next()
		},
	})(c)
}
