package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Unverified" {
		t := views.SignUpPage()
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Unverified" {
		data := types.SignupData{
			Fullname:  r.FormValue("fullname"),
			Username:  r.FormValue("username"),
			Password:  r.FormValue("pswd"),
			CPassword: r.FormValue("cpswd"),
			IsAdmin:   false,
		}
		err := models.SignUp(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
