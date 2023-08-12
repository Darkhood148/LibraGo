package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func BookList(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		db, err := models.Connection()
		if err != nil {
			data, err := models.FetchBooks(db)
			if err != nil {
				renderErrorPage(w, err.Error())
			} else {
				t := views.RenderPage("bookList.html")
				t.Execute(w, data)
			}
		} else {
			renderErrorPage(w, "Cannot Connect to db")
		}
	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}
