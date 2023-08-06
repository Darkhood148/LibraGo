package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func CheckRequest(w http.ResponseWriter, r *http.Request) {
	data := models.FetchRequests()
	fmt.Println("data", data)
	t := views.CheckRequestPage()
	t.Execute(w, data)
}

func CheckRequestPost(w http.ResponseWriter, r *http.Request) {
	temp := r.FormValue("actionInfo")
	fmt.Println(temp)
	words := strings.Split(temp, "-")
	var status string
	if words[1] == "a" {
		status = "approve"
	} else {
		status = "denied"
	}
	checkoutid, err := strconv.Atoi(words[0])
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		models.CheckRequest(checkoutid, status)
	}
}
