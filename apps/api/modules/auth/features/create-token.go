package features

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CreateJwtTokenDTO struct {
	UserId int
}

func CreateJwtToken(dto CreateJwtTokenDTO) (string, error) {
	expiration := time.Now().Add(72 * time.Hour)

	claims := jwt.MapClaims{
		"id":  dto.UserId,
		"exp": expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
