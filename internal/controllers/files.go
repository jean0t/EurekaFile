package controllers

import (
	"net/http"
	"html/template"
	"path/filepath"
)

const basePath string = "internal/views"

var templ = template.Must(template.ParseFiles(
	filepath.Join(basePath, "files.tmpl"),
	filepath.Join(basePath, "upload.tmpl"),
	filepath.Join(basePath, "navbar.tmpl"),
))

func Files(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Files", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "Upload", nil)
}
