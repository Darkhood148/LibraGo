package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"net/http"
)

func ReturnBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		data := r.FormValue("actionInfo")
		err := models.ReturnBook(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(types.Success))
		}
	} else {
		w.Write([]byte(types.NotClient))
	}
}
