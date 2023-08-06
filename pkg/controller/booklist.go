package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func BookList(w http.ResponseWriter, r *http.Request) {
	data := models.FetchBooks()
	t := views.BookListPage()
	t.Execute(w, data)
}