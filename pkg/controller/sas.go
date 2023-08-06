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
	uname := middleware.VerifyJWT(w, r)
	if uname != "" {
		if middleware.VerifyAdmin(uname) {
			t := views.SasPage()
			t.Execute(w, nil)
		}
	} else {
		w.Write([]byte("Login Please"))
	}
}

func SasPost(w http.ResponseWriter, r *http.Request) {
	val1, _ := strconv.Atoi(r.FormValue("bookid"))
	val2, _ := strconv.Atoi(r.FormValue("quantity"))
	data := types.SasData{
		Bookid:   val1,
		Option:   r.FormValue("options"),
		Quantity: val2,
	}
	models.Sas(data)
}
