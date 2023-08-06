package controller

import (
	"fmt"
	"mvc/pkg/models"
	"net/http"
)

func ReturnDeniedBookPost(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("actionInfo")
	fmt.Println("data", data)
	models.ReturnDeniedBook(data)
}
