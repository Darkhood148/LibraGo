package models

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims2 struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKeyy = []byte(os.Getenv("DB_JWTSECRET"))

func Logout() (http.Cookie, error) {
	expirationTime := time.Now()
	claims := &Claims2{
		Username: "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeyy)
	if err != nil {
		return http.Cookie{}, err
	}
	return http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}, nil
}
