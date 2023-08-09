package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func IssueBook(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) != "Unverified" {
		t := views.IssueBookPage()
		t.Execute(w, nil)
	} else {
		w.Write([]byte("You need to be logged in to access this."))
	}
}

func IssueBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) != "Unverified" {
		val, err := strconv.Atoi(r.FormValue("bookid"))
		if err != nil {
			w.Write([]byte("Error Occured while parsing input values"))
		} else {
			data := types.IssueBookData{
				Bookid:   val,
				Username: middleware.VerifyJWT(w, r),
			}
			err := models.IssueBook(data)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte("Success"))
			}
		}
	} else {
		w.Write([]byte("Please login to issue a book"))
	}
}
