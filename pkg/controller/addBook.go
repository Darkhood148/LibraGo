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
		renderAddPage(w, "0", "")
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func AddBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		quant, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			renderAddPage(w, "2", err.Error())
		} else if quant <= 0 {
			renderAddPage(w, "2", "Please enter positive values in field quantity")
		} else {
			data := types.Book{
				Bookname: r.FormValue("bookname"),
				Author:   r.FormValue("author"),
				Quantity: quant,
			}
			err := models.AddBook(data)
			if err != nil {
				renderAddPage(w, "2", err.Error())
			} else {
				renderAddPage(w, "1", "")
			}
		}
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func renderAddPage(w http.ResponseWriter, status string, errMess string) {
	t := views.AddBookPage()
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
