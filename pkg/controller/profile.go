package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	uname := middleware.VerifyJWT(w, r)
	if uname != "" {
		t := views.ProfilePage()
		info := types.LoginData{
			Username: uname,
		}
		t.Execute(w, info)
	} else {
		w.Write([]byte("Login you fucking idiot"))
	}
}
