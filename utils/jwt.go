package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret" // DEMO key

func GenerateToken(email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return fmt.Errorf("could not parse token. Error message: %s", err)
	}

	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return errors.New("token is invalid")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("token claims are invalid")
	// }

	// email := claims["email"].(string)
	// userID := claims["user_id"].(int64)

	return nil
}
