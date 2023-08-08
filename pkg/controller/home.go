package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/views"
	"net/http"
)

type Data struct {
	Status bool `json:"status"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := views.StartPage()
	var data Data
	if middleware.TypeOfUser(w, r) == "Unverified" {
		data.Status = false
	} else {
		data.Status = true
	}
	t.Execute(w, data)
}
