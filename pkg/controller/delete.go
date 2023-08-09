package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		t := views.DeleteBookPage()
		t.Execute(w, nil)
	} else {
		w.Write([]byte("You need to be an admin to access this."))
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		data, err := strconv.Atoi(r.FormValue("bookid"))
		if err != nil {
			w.Write([]byte("Error Occured while parsing input values"))
		} else {
			err := models.DeleteBook(data)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte("Success"))
			}
		}
	} else {
		w.Write([]byte("You need to be an admin to access this."))
	}
}
