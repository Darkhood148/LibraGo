package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		renderDeletePage(w, "0", "")
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		data, err := strconv.Atoi(r.FormValue("bookid"))
		if err != nil {
			renderDeletePage(w, "2", err.Error())
		} else {
			err := models.DeleteBook(data)
			if err != nil {
				renderDeletePage(w, "2", err.Error())
			} else {
				renderDeletePage(w, "1", "")
			}
		}
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func renderDeletePage(w http.ResponseWriter, status string, errMess string) {
	t := views.DeleteBookPage()
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
