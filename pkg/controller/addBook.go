package controller

import (
	"fmt"
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	uname := middleware.VerifyJWT(w, r)
	if uname != "" {
		if middleware.VerifyAdmin(uname) {
			t := views.AddBookPage()
			t.Execute(w, nil)
		} else {
			fmt.Println("Not Admin")
		}
	} else {
		fmt.Println("Please Login")
	}
}

func AddBookPost(w http.ResponseWriter, r *http.Request) {
	quant, _ := strconv.Atoi(r.FormValue("quantity"))
	data := types.Book{
		Bookname: r.FormValue("bookname"),
		Author:   r.FormValue("author"),
		Quantity: quant,
	}
	fmt.Println(data)
	models.AddBook(data)
}
