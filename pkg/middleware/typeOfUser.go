package middleware

import (
	"fmt"
	"net/http"
)

func TypeOfUser(w http.ResponseWriter, r *http.Request) string {
	uname := VerifyJWT(w, r)
	if uname != "" {
		fmt.Println("uname", uname)
		if VerifyAdmin(uname) {
			return "Admin"
		} else {
			return "User"
		}
	} else {
		return "Unverified"
	}
}