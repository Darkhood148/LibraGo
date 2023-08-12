package controller

import (
	"errors"
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func CheckRequest(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		data, err := models.FetchRequests()
		if err != nil {
			renderErrorPage(w, err.Error())
		} else {
			t := views.CheckRequestPage()
			t.Execute(w, data)
		}
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}

func CheckRequestPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Admin {
		temp := r.FormValue("actionInfo")
		words := strings.Split(temp, "-")
		var status string
		if words[1] == "a" {
			status = "approve"
		} else if words[1] == "d" {
			status = "denied"
		} else {
			renderErrorPage(w, "Invalid Request")
			panic(errors.New("wrong input"))
		}
		checkoutid, err := strconv.Atoi(words[0])
		if err != nil {
			renderErrorPage(w, "Error occured while parsing input")
		} else {
			err := models.CheckRequest(checkoutid, status)
			if err != nil {
				renderErrorPage(w, err.Error())
			} else {
				http.Redirect(w, r, "/checkRequest", http.StatusSeeOther)
			}
		}
	} else {
		renderInvalidPage(w, string(types.NotAdmin))
	}
}
