package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	t := views.SignUpPage()
	t.Execute(w, nil)
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	data := types.SignupData{
		Fullname:  r.FormValue("fullname"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("pswd"),
		CPassword: r.FormValue("cpswd"),
		IsAdmin:   false,
	}
	fmt.Println(data)
	models.SignUp(data)
}
