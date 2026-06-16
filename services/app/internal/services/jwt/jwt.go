package jwt

import (
	"app/config/settings"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userID uint) (string, error) {
	claims :=  jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(settings.Envs.JWT_SECRET))
}