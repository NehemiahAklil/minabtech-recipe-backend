package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey []byte

func GenerateToken(user entity.User, expiresIn time.Duration) (string, error) {

	if SecretKey != nil {
		return "", errors.New("JWT Seceret not set")
	}
	claims := jwt.MapClaims{
		"exp": time.Now().Add(expiresIn).Unix(),
		"iat": time.Now().Unix(),
		"uid": user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("error generating token : %w", err)
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return nil, nil
}
