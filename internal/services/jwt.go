package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(getJWTSecret())
}
