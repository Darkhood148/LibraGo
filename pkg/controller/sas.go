package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func Sas(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		t := views.SasPage()
		t.Execute(w, nil)
	} else {
		w.Write([]byte("You need to be an admin to access this."))
	}
}

func SasPost(w http.ResponseWriter, r *http.Request) {
	val1, err1 := strconv.Atoi(r.FormValue("bookid"))
	val2, err2 := strconv.Atoi(r.FormValue("quantity"))
	if err1 != nil && err2 != nil {
		w.Write([]byte("Error Occured"))
	} else {
		data := types.SasData{
			Bookid:   val1,
			Option:   r.FormValue("options"),
			Quantity: val2,
		}
		err := models.Sas(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Success"))
		}
	}
}
