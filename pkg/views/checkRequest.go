package views

import (
	"html/template"
)

func CheckRequestPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/checkRequests.html"))
	return temp
}
