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
		t := views.DeleteBookPage()
		t.Execute(w, nil)
	} else {
		w.Write([]byte(types.NotAdmin))
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		data, err := strconv.Atoi(r.FormValue("bookid"))
		if err != nil {
			w.Write([]byte("Error Occured while parsing input values"))
		} else {
			err := models.DeleteBook(data)
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
