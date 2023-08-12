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
		bookid, err1 := strconv.Atoi(r.FormValue("bookid"))
		quantity, err2 := strconv.Atoi(r.FormValue("quantity"))
		if err1 != nil || err2 != nil {
			renderSasPage(w, "2", "Error occured while parsing input values")
		} else if quantity <= 0 || bookid <= 0 {
			renderSasPage(w, "2", "Please enter input values")
		} else {
			data := types.SasData{
				Bookid:   bookid,
				Option:   r.FormValue("options"),
				Quantity: quantity,
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
	t := views.RenderPage("sas.html")
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
