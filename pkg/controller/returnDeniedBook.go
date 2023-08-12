package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"net/http"
)

func ReturnDeniedBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		data := r.FormValue("actionInfo")
		err := models.ReturnDeniedBook(data)
		if err != nil {
			renderErrorPage(w, err.Error())
		} else {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}
