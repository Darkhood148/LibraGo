package views

import (
	"html/template"
)

func RequestAdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/requestAdmin.html"))
	return temp
}
