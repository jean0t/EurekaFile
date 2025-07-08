package middleware

import (
	"net/http"
	"time"
	"fmt"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s]: serving %s to %s\n", time.Now().Format("02/01/2006 15:04:05"), r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
