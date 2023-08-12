package views

import (
	"html/template"
	"path/filepath"
)

func RenderPage(file string) *template.Template {
	temp := template.Must(template.ParseFiles(filepath.Join("templates", file)))
	return temp
}
