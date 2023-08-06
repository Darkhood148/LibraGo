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
	r.HandleFunc("/addbook", controller.AddBook).Methods("GET")
	r.HandleFunc("/addbook", controller.AddBookPost).Methods("POST")
	r.HandleFunc("/booklist", controller.BookList).Methods("GET")
	r.HandleFunc("/sas", controller.Sas).Methods("GET")
	r.HandleFunc("/sas", controller.SasPost).Methods("POST")
	r.HandleFunc("/delete", controller.Delete).Methods("GET")
	r.HandleFunc("/delete", controller.DeletePost).Methods("POST")
	r.HandleFunc("/issuebook", controller.IssueBook).Methods("GET")
	r.HandleFunc("/issuebook", controller.IssueBookPost).Methods("POST")
	r.HandleFunc("/checkrequest", controller.CheckRequest).Methods("GET")
	r.HandleFunc("/checkrequest", controller.CheckRequestPost).Methods("POST")
	r.HandleFunc("/returnbook", controller.ReturnBookPost).Methods("POST")
	r.HandleFunc("/returndeniedbook", controller.ReturnDeniedBookPost).Methods("POST")

	http.ListenAndServe(":8000", r)
}
