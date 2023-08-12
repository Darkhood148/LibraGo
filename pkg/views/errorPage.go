package views

import (
	"html/template"
)

func ErrorPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/err404.html"))
	return temp
}
