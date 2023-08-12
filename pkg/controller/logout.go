package controller

import (
	"mvc/pkg/models"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := models.Logout()
	if err != nil {
		renderErrorPage(w, err.Error())
	} else {
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
