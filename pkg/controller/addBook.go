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
	if middleware.TypeOfUser(w, r) == types.Admin {
		t := views.AddBookPage()
		t.Execute(w, nil)
	} else {
		w.Write([]byte(types.NotAdmin))
	}
}

func AddBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		quant, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			w.Write([]byte(err.Error()))
		} else if quant <= 0 {
			w.Write([]byte("Please enter positive values"))
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
				w.Write([]byte(types.Success))
			}
		}
	} else {
		w.Write([]byte(types.NotAdmin))
	}
}
