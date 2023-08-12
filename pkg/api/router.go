package api

import (
	"net/http"

	"mvc/pkg/controller"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/signup", controller.Signup).Methods("GET")
	r.HandleFunc("/signup", controller.SignupPost).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/profile", controller.Profile).Methods("GET")
	r.HandleFunc("/addBook", controller.AddBook).Methods("GET")
	r.HandleFunc("/addBook", controller.AddBookPost).Methods("POST")
	r.HandleFunc("/bookList", controller.BookList).Methods("GET")
	r.HandleFunc("/sas", controller.Sas).Methods("GET")
	r.HandleFunc("/sas", controller.SasPost).Methods("POST")
	r.HandleFunc("/delete", controller.Delete).Methods("GET")
	r.HandleFunc("/delete", controller.DeletePost).Methods("POST")
	r.HandleFunc("/issueBook", controller.IssueBook).Methods("GET")
	r.HandleFunc("/issueBook", controller.IssueBookPost).Methods("POST")
	r.HandleFunc("/checkRequest", controller.CheckRequest).Methods("GET")
	r.HandleFunc("/checkRequest", controller.CheckRequestPost).Methods("POST")
	r.HandleFunc("/returnBook", controller.ReturnBookPost).Methods("POST")
	r.HandleFunc("/returnDeniedBook", controller.ReturnDeniedBookPost).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/requestAdmin", controller.GetAdminAccess).Methods("GET")
	r.HandleFunc("/requestAdmin", controller.GetAdminAccessPost).Methods("POST")

	http.ListenAndServe(":8000", r)
}
