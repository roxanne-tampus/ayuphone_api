package utils

import (
	"ayuphone_api/internal/models"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims defines the JWT claims
type Claims struct {
	ID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the given user
func GenerateToken(user *models.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY") // Ensure correct environment variable name

	claims := Claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken parses and validates the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY")) // Ensure correct environment variable name

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
