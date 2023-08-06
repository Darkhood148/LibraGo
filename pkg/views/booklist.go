package views

import (
	"html/template"
)

func BookListPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/bookList.html"))
	return temp
}