package utils

import (
	"errors"
	"time"
	"user-management-api/internal/models"

	"github.com/dgrijalva/jwt-go"
)

const JWT_SECRET = "my_secret_key"

// GenerateToken generates JWT token
func GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies JWT token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
