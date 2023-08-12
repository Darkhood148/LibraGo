package controller

import (
	"mvc/pkg/views"
	"net/http"
)

type err struct {
	Reason string `json:"string"`
}

func renderErrorPage(w http.ResponseWriter, reason string) {
	t := views.RenderPage("err404.html")
	info := err{
		Reason: reason,
	}
	t.Execute(w, info)
}
