package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/khairulharu/gojwt/domain"
)

var Key = []byte("mysecretkey")

type MyCustomClaims struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

func CreateToken(user *domain.User) (string, error) {
	claims := MyCustomClaims{
		user.Username,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(Key)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("error token validate")
	}

	return claims, nil
}
