package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) != types.Unverified {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		t := views.LoginPage()
		t.Execute(w, nil)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Unverified {
		data := types.LoginData{
			Username: r.FormValue("username"),
			Password: r.FormValue("pswd"),
		}
		cookie, err := models.Login(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
