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

func IssueBook(w http.ResponseWriter, r *http.Request) {
	t := views.IssueBookPage()
	t.Execute(w, nil)
}

func IssueBookPost(w http.ResponseWriter, r *http.Request) {
	uname := middleware.VerifyJWT(w, r)
	if uname != "" {
		fmt.Println(uname)
		val, err := strconv.Atoi(r.FormValue("bookid"))
		if err != nil {
			fmt.Println("Error Occured")
		} else {
			data := types.IssueBookData{
				Bookid:   val,
				Username: uname,
			}
			models.IssueBook(data)
		}
	} else {
		fmt.Println("Not Logged In")
	}
}
