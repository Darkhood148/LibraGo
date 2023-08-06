package views

import (
	"html/template"
)

func ProfilePage(admin bool) *template.Template {
	if admin {
		temp := template.Must(template.ParseFiles("templates/profileAdmin.html"))
		return temp
	} else {
		temp := template.Must(template.ParseFiles("templates/profile.html"))
		return temp
	}
}
