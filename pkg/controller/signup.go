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
		w.Write([]byte("Error Occured"))
	}
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	data := types.SignupData{
		Fullname:  r.FormValue("fullname"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("pswd"),
		CPassword: r.FormValue("cpswd"),
		IsAdmin:   false,
	}
	models.SignUp(data)
}
