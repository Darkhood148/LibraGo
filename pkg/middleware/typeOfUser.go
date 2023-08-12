package middleware

import (
	"mvc/pkg/types"
	"net/http"
)

func TypeOfUser(w http.ResponseWriter, r *http.Request) types.User {
	username := VerifyJWT(w, r)
	if username != "" {
		if VerifyAdmin(username) {
			return types.Admin
		} else {
			return types.Client
		}
	} else {
		return types.Unverified
	}
}
