package models

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/yaml.v3"
)

var jwtKeyy []byte

type JWTConfigg struct {
	JWT_SECRET string `yaml:"JWT_SECRET"`
}

func JwtSecretKeyy() {
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

	jwtKeyy = []byte(config.JWT_SECRET)
}

type Claims2 struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Logout() (http.Cookie, error) {
	expirationTime := time.Now()
	claims := &Claims2{
		Username: "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JwtSecretKeyy()
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
