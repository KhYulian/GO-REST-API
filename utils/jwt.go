package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	const secretKey = "secret" // DEMO key

	return token.SignedString([]byte(secretKey))
}
