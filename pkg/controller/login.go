package controller

import (
	"fmt"
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) != "Unverified" {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		t := views.LoginPage()
		t.Execute(w, nil)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(middleware.TypeOfUser(w, r))
	if middleware.TypeOfUser(w, r) == "Unverified" {
		data := types.LoginData{
			Username: r.FormValue("username"),
			Password: r.FormValue("pswd"),
		}
		cookie, err := models.Login(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			http.SetCookie(w, &cookie)
			w.Write([]byte("Success"))
		}
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
