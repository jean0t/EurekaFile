package middleware

import (
	"net/http"
	
	"github.com/jean0t/EurekaFile/internal/auth"
	"github.com/golang-jwt/jwt/v5"
)


func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		var tokenStr string = cookie.Value
		var claims *auth.Claims = &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
