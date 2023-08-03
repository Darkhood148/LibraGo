package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	t := views.LoginPage()
	t.Execute(w, nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	data := types.LoginData{
		Username: r.FormValue("username"),
		Password: r.FormValue("pswd"),
	}
	fmt.Println(data)
	models.Login(data, w, r)
}