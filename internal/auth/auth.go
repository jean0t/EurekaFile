package auth

import (
	"time"
	"os"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)


type contextKey string
const UserContextKey = contextKey("user")


type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}


func CreateToken(username string, expiration time.Time) string {
	var signedToken string
	var err error
	var claims *Claims = &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	signedToken, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ""
	}

	return signedToken
}


func GetUser(r *http.Request) *Claims {
	var (
		claims *Claims
		ok bool
	)

	claims, ok = r.Context().Value(UserContextKey).(*Claims)
	if !ok {
		return nil
	}

	return claims
}
