package controllers

import (
	"net/http"
	"html/template"
	"path/filepath"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var templ = template.Must(template.ParseFiles(
		filepath.Join("internal", "views", "login.tmpl"),
	))

	_ = templ.ExecuteTemplate(w, "Login", nil)
}
