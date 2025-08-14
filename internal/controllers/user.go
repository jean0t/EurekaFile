package controllers

import (
	"net/http"
	"strings"
	"time"
	"fmt"

	"github.com/jean0t/EurekaFile/internal/auth"
	"github.com/jean0t/EurekaFile/internal/database"

	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var username string = strings.TrimSpace(r.FormValue("username"))
	var password string = strings.TrimSpace(r.FormValue("password"))
	var db *gorm.DB


	if len(strings.Split(username, "")) < 3 || len(strings.Split(password, "")) < 3 {
                http.Redirect(w, r, "/", http.StatusUnauthorized)
                return
        }


	db, err = database.ConnectToDB()
	if err != nil {
		fmt.Println("[!] Error connecting to DB")
		http.Error(w, "<h1>There was an unexpected error on the server</h1>", http.StatusInternalServerError)
		return
	}

	err = database.IsValidUser(db, username, password)
	if err != nil {
		fmt.Println("Error validating user")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Failed to authenticate</h1>")
		return
	}
	
	var expiration time.Time = time.Now().Add(24*time.Hour)
	var signedToken string = auth.CreateToken(username, expiration)

	http.SetCookie(w, &http.Cookie {
		Name: "Authentication",
		Value: signedToken,
		Path: "/",
		HttpOnly: true,
		Expires: expiration,
	})

	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "Authentication",
		Value: "",
		Path: "/",
		HttpOnly: true,
		Expires: time.Unix(0, 0),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
