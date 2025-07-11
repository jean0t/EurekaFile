package router

import (
	"net/http"

	"github.com/jean0t/EurekaFile/internal/controllers"
)

func GetRouter() *http.ServeMux {
	var Router *http.ServeMux = http.NewServeMux()
	
	// routes
	Router.HandleFunc("/", controllers.Index)
	Router.HandleFunc("/upload", controllers.Upload)
	Router.HandleFunc("/files", controllers.Files)
	Router.HandleFunc("/login", controllers.Login)
	Router.HandleFunc("/logout", controllers.Logout)

	return Router
}
