package views

import (
	"html/template"
)

func InvalidPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/invalidAccess.html"))
	return temp
}
