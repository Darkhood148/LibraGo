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
		renderLoginPage(w, "0", "")
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
			renderLoginPage(w, "2", err.Error())
		} else {
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}

func renderLoginPage(w http.ResponseWriter, status string, errMess string) {
	t := views.LoginPage()
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
