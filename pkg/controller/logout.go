package controller

import (
	"mvc/pkg/models"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := models.Logout()
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
