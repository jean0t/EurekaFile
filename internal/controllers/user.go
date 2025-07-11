package controllers

import (
	"net/http"
	"github.com/jean0t/EurekaFile/internal/auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	var err error

	cookie, err = r.Cookie("auth")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
