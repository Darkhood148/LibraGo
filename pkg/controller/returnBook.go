package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"net/http"
)

func ReturnBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) != "Unverified" {
		data := r.FormValue("actionInfo")
		err := models.ReturnBook(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Success"))
		}
	} else {
		w.Write([]byte("You need yo be logged in to access this."))
	}
}
