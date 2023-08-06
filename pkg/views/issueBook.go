package views

import (
	"html/template"
)

func IssueBookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/issueBook.html"))
	return temp
}
