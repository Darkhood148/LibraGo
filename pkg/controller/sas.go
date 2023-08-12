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
	if middleware.TypeOfUser(w, r) == types.Admin {
		renderSasPage(w, "0", "")
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func SasPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		val1, err1 := strconv.Atoi(r.FormValue("bookid"))
		val2, err2 := strconv.Atoi(r.FormValue("quantity"))
		if err1 != nil || err2 != nil {
			renderSasPage(w, "2", "Error occured while parsing input values")
		} else if val2 <= 0 || val1 <= 0 {
			renderSasPage(w, "2", "Please enter input values")
		} else {
			data := types.SasData{
				Bookid:   val1,
				Option:   r.FormValue("options"),
				Quantity: val2,
			}
			err := models.Sas(data)
			if err != nil {
				renderSasPage(w, "2", err.Error())
			} else {
				renderSasPage(w, "1", "")
			}
		}
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func renderSasPage(w http.ResponseWriter, status string, errMess string) {
	t := views.SasPage()
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
