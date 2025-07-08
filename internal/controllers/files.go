package controllers

import (
	"net/http"
	"html/template"
	"path/filepath"
)

var templ = template.Must(template.ParseFiles(
	filepath.Join("views", "files.tmpl"),
	filepath.Join("views", "upload.tmpl"),
	filepath.Join("views", "navbar.tmpl"),
))

func Files(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Files", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Upload", nil)
}
