package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TOKEN_DURATION_IN_HOURS = 72

type JwtUtils struct{}

func NewJwtUtils() JwtUtils {
	return JwtUtils{}
}

func (this *JwtUtils) GenerateToken(id uint) (string, error) {
	expiration := time.Now().Add(TOKEN_DURATION_IN_HOURS * time.Hour)
	claims := jwt.MapClaims{
		"id":  id,
		"exp": expiration.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
