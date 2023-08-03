package api

import (
	"net/http"

	"mvc/pkg/controller"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/signup", controller.Signup).Methods("GET")
	r.HandleFunc("/signup", controller.SignupPost).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/profile", controller.Profile).Methods("GET")

	http.ListenAndServe(":8000", r)
}
