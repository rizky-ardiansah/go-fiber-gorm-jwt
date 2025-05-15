package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, email string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	jwtExpiresInStr := os.Getenv("JWT_EXPIRES_IN")

	duration, err := time.ParseDuration(jwtExpiresInStr)
	if err != nil {
		return "", err
	}

	claims := MyCustomClaims{
		userID,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",                               // opsional
			Subject:   strconv.FormatUint(uint64(userID), 10), // opsional
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return ss, nil

}

func ValidateJWT(tokenString string) (*MyCustomClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
