package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/yaml.v3"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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

func VerifyJWT(w http.ResponseWriter, r *http.Request) string {
	JwtSecretKey()
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}

	tknStr := c.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return ""
	}
	return claims.Username
}
