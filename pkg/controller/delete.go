package controller

import (
	"fmt"
	"mvc/pkg/middleware"
	"mvc/pkg/views"
	"mvc/pkg/models"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	uname := middleware.VerifyJWT(w, r)
	if uname != "" {
		if middleware.VerifyAdmin(uname) {
			t := views.DeleteBookPage()
			t.Execute(w, nil)
		}
	} else {
		w.Write([]byte("Login Please"))
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	data, err := strconv.Atoi(r.FormValue("bookid"))
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		models.DeleteBook(data)
	}
}
