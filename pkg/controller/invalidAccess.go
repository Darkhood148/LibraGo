package controller

import (
	"mvc/pkg/views"
	"net/http"
)

type data struct {
	Reason string `json:"string"`
}

func renderInvalidPage(w http.ResponseWriter, reason string) {
	t := views.RenderPage("invalidAccess.html")
	info := data{
		Reason: reason,
	}
	t.Execute(w, info)
}
