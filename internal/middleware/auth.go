package middleware

import (
	"net/http"
	"os"
	"context"

	"github.com/jean0t/EurekaFile/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey = contextKey("user")

func WithAuth(next http.Handler) http.Handler {
	var jwtKey []byte = []byte(os.Getenv("JWT_SECRET"))

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var (
			cookie *http.Cookie
			err error
		)

		cookie, err = r.Cookie("auth")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
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

		var ctx context.Context = context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func GetUser(r *http.Request) *auth.Claims {
	claims, ok := r.Context().Value(userContextKey).(*auth.Claims)
	if !ok {
		return nil
	}

	return claims
}
