package middleware

import (
	"net/http"
	"strings"
	"time"
	"fmt"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ip, port string
		ip, port, _= strings.Cut(r.RemoteAddr, ":")
		fmt.Printf("[%s]: serving %s to ip %s through port %s\n", time.Now().Format("02/01/2006 15:04:05"), r.RequestURI, ip, port)
		next.ServeHTTP(w, r)
	})
}
