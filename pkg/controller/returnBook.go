package controller

import (
	"mvc/pkg/models"
	"net/http"
)

func ReturnBookPost(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("actionInfo")
	models.ReturnBook(data)
}