package models

import (
	"errors"
	"log"
	"mvc/pkg/types"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

var jwtKey []byte

type JWTConfig struct {
	JWT_SECRET string `yaml:"JWT_SECRET"`
}

func JwtSecretKey() {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config JWTConfig
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	jwtKey = []byte(config.JWT_SECRET)
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(data types.LoginData) (http.Cookie, error) {
	db, err := Connection()
	if err != nil {
		return http.Cookie{}, err
	}
	defer db.Close()
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
					expirationTime := time.Now().Add(5 * time.Hour)
					claims := &Claims{
						Username: data.Username,
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(expirationTime),
						},
					}
					JwtSecretKey()
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
