package models

import (
	"errors"
	"mvc/pkg/types"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(os.Getenv("DB_JWTSECRET"))

func Login(data types.LoginData) (http.Cookie, error) {
	db, err := Connection()
	if err != nil {
		return http.Cookie{}, err
	}
	check := "SELECT * FROM users WHERE username=(?)"
	res, err := db.Query(check, data.Username)
	if err != nil {
		return http.Cookie{}, err
	} else if !res.Next() {
		return http.Cookie{}, errors.New("username does not exists")
	} else {
		pswd := []byte(data.Password)
		query := "SELECT hash FROM users WHERE username=(?)"
		res, err := db.Query(query, data.Username)
		if err != nil {
			return http.Cookie{}, err
		} else {
			var hashedPassword string
			res.Next()
			err := res.Scan(&hashedPassword)
			if err != nil {
				return http.Cookie{}, err
			} else {
				err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), pswd)
				if err != nil {
					return http.Cookie{}, errors.New("incorrect Password")
				} else {
					expirationTime := time.Now().Add(5 * time.Minute)
					claims := &Claims{
						Username: data.Username,
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(expirationTime),
						},
					}
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
					tokenString, err := token.SignedString(jwtKey)
					if err != nil {
						return http.Cookie{}, err
					}
					return http.Cookie{
						Name:    "token",
						Value:   tokenString,
						Expires: expirationTime,
					}, nil
				}
			}
		}
	}
}
