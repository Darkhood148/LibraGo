package middleware

import (
	"mvc/pkg/types"
	"net/http"
)

func TypeOfUser(w http.ResponseWriter, r *http.Request) types.User {
	uname := VerifyJWT(w, r)
	if uname != "" {
		if VerifyAdmin(uname) {
			return types.Admin
		} else {
			return types.Client
		}
	} else {
		return types.Unverified
	}
}
