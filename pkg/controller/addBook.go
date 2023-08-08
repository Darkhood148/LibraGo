package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		t := views.AddBookPage()
		t.Execute(w, nil)
	} else {
		w.Write([]byte("You need to be an admin to access this."))
	}
}

func AddBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		quant, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			data := types.Book{
				Bookname: r.FormValue("bookname"),
				Author:   r.FormValue("author"),
				Quantity: quant,
			}
			err := models.AddBook(data)
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
