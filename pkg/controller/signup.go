package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	t := views.SignUpPage()
	t.Execute(w, nil)
}