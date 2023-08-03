package models

import (
	"fmt"
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

func Login(data types.LoginData, w http.ResponseWriter, r *http.Request) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	} else {
		check := "SELECT * FROM users WHERE username=(?)"
		res, err := db.Query(check, data.Username)
		if err != nil {
			fmt.Println("Error Occurred")
		} else if !res.Next() {
			fmt.Println("User does not exist")
		} else {
			pswd := []byte(data.Password)
			query := "SELECT hash FROM users WHERE username=(?)"
			res, err := db.Query(query, data.Username)
			if err != nil {
				fmt.Println("Error Occurred", err)
			} else {
				var hashedPassword string
				res.Next()
				err := res.Scan(&hashedPassword)
				if err != nil {
					fmt.Println("Error Occurred", err)
				} else {
					err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), pswd)
					if err != nil {
						fmt.Println("Incorrect")
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
							w.WriteHeader(http.StatusInternalServerError)
							return
						}
						http.SetCookie(w, &http.Cookie{
							Name:    "token",
							Value:   tokenString,
							Expires: expirationTime,
						})
						fmt.Println(tokenString)
					}
				}
			}
		}
	}
}
