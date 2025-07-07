package router

import (
	"net/http"
	"fmt"

//	"github.com/JMFern01/EurekaFile/internal/controllers"
)

func GetRouter() *http.ServeMux {
	var Router *http.ServeMux = http.NewServeMux()
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World</h1>")
	})
/*
	Router.HandleFunc("/", controllers.Index)
	Router.HandleFunc("/upload", controllers.Index)
	Router.HandleFunc("/files", controllers.Files)
	Router.HandleFunc("/login", controllers.Login)
	Router.HandleFunc("/logout", controllers.Logout)
*/
	return Router
}
