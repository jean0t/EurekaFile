package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	
	"github.com/JMFern01/EurekaFile/internal/router"
	"github.com/JMFern01/EurekaFile/internal/middleware"
)

func main() {
	var startServer *bool = flag.Bool("s", false, "Starts the server")
    var portServer *int = flag.Int("p", 8000, "Port to link the server")
	flag.Parse()

	if *startServer {
		var sigChan chan os.Signal = make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		fmt.Println("[*] Starting Server...")
		go func() {
			var Router = router.GetRouter()
			var loggedRouter http.Handler = middleware.LoggingMiddleware(Router)

			err := http.ListenAndServe(fmt.Sprintf(":%d", *portServer), loggedRouter)
			if err != nil {
				fmt.Println("Error while starting the server")
			}
		}()
		fmt.Println("[*] Server Started.")
        fmt.Printf("[*] You can access the server in http://localhost:%d/\n", *portServer)

		<-sigChan
		fmt.Println("\n[*] Closing Server.")
	}
	
}
