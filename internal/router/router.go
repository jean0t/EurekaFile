package router

import (
	"net/http"

	"github.com/jean0t/EurekaFile/internal/controllers"
	"github.com/jean0t/EurekaFile/internal/middleware"
)

func GetRouter() *http.ServeMux {
	var Router *http.ServeMux = http.NewServeMux()
	
	// routes
	Router.HandleFunc("/", controllers.Index)
	Router.Handle("/upload", middleware.WithAuth(http.HandlerFunc(controllers.Upload)))
	Router.Handle("/files", middleware.WithAuth(http.HandlerFunc(controllers.Files)))
	Router.Handle("/files/", middleware.WithAuth(http.HandlerFunc(controllers.DownloadFile)))
	Router.Handle("/login", http.HandlerFunc(controllers.Login))
	Router.Handle("/logout", middleware.WithAuth(http.HandlerFunc(controllers.Logout)))

	return Router
}
