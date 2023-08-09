package controller

import (
	"errors"
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func CheckRequest(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		data, err := models.FetchRequests()
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			t := views.CheckRequestPage()
			t.Execute(w, data)
		}
	} else {
		w.Write([]byte("You need to be an admin to access this."))
	}
}

func CheckRequestPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == "Admin" {
		temp := r.FormValue("actionInfo")
		words := strings.Split(temp, "-")
		var status string
		if words[1] == "a" {
			status = "approve"
		} else if words[1] == "d" {
			status = "denied"
		} else {
			w.Write([]byte("Invalid Request"))
			panic(errors.New("wrong input"))
		}
		checkoutid, err := strconv.Atoi(words[0])
		if err != nil {
			w.Write([]byte("Error Occured while parsing input"))
		} else {
			err := models.CheckRequest(checkoutid, status)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte("Success"))
			}
		}
	} else {
		w.Write([]byte("You need to be an admin to access this."))
	}
}
