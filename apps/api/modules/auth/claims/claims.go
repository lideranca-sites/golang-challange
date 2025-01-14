package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.Claims
	UserId int `json:"user_id"`
}
