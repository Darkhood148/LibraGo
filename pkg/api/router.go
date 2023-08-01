package api

import (
	"net/http"

	"mvc/pkg/controller"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("GET")

	http.ListenAndServe(":8000", r)
}