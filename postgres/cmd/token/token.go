package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

const (
	Duration = 24 * time.Hour
)

func Generate(userID string) (string, error) {
	expirationTime := time.Now().Add(Duration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Создаем токен с подписью HMAC-SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func Validate(tokenString string) (bool, string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return false, "", fmt.Errorf("failed to parse token: %v", err)
	}
	if !token.Valid {
		return false, "", fmt.Errorf("invalid token")
	}

	// Проверка срока действия токена
	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return false, "", fmt.Errorf("token has expired")
	}

	return true, claims.UserID, nil
}
