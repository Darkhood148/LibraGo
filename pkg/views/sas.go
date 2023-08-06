package views

import (
	"html/template"
)

func SasPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/sas.html"))
	return temp
}
