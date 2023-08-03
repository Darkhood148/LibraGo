package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(os.Getenv("DB_JWTSECRET"))

func VerifyJWT(w http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
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
