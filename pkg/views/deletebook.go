package views

import (
	"html/template"
)

func DeleteBookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/delete.html"))
	return temp
}
