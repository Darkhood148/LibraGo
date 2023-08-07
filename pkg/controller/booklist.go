package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func BookList(w http.ResponseWriter, r *http.Request) {
	data, err := models.FetchBooks()
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		t := views.BookListPage()
		t.Execute(w, data)
	}
}
