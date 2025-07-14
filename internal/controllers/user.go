package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/jean0t/EurekaFile/internal/auth"
	"github.com/jean0t/EurekaFile/internal/database"

	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var username string = r.FormValue("username")
	var password string = r.FormValue("password")
	var db *gorm.DB

	db, err= database.ConnectToDB()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}

	err = database.IsValidUser(db, username, strings.TrimSpace(password))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	
	var expiration time.Time = time.Now().Add(24*time.Hour)
	var signedToken string = auth.CreateToken(username, expiration)

	http.SetCookie(w, &http.Cookie {
		Name: "auth",
		Value: signedToken,
		Path: "/",
		HttpOnly: true,
		Secure: true,
		Expires: expiration,
	})

	http.Redirect(w, r, "/upload", http.StatusSeeOther)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "auth",
		Value: "",
		Path: "/",
		Expires: time.Unix(0, 0),
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
