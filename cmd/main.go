package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	
	"github.com/JMFern01/EurekaFile/internal/router"
)

func main() {
	var startServer *bool = flag.Bool("s", false, "Starts the server")
	flag.Parse()

	if *startServer {
		var sigChan chan os.Signal = make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		fmt.Println("[*] Starting Server...")
		go func() {
			var Router = router.GetRouter()
			err := http.ListenAndServe(":8000", Router)
			if err != nil {
				fmt.Println("Error while starting the server")
			}
		}()
		fmt.Println("[*] Server Started.")

		<-sigChan
		fmt.Println("\n[*] Closing Server.")
	}
	
}
