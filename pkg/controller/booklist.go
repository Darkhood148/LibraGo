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
		data, err := models.FetchBooks()
		if err != nil {
			renderErrorPage(w, err.Error())
		} else {
			t := views.BookListPage()
			t.Execute(w, data)
		}
	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}
