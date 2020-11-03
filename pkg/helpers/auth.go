package helpers

import (
	"fmt"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var MySigningKey = "yJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik1rWkNOMEk1UXpGRk4wUTFRVUl4TkVFNU1qUkNPRGhETlRrMk9EVTFOakl4TlRjM01qbEJNUSJ9.eyJpc3MiOiJodHRwczovL2NoYWxrYW4uYXV0aDAuY29tLyIsInN1YiI6IkhZT2VaZGtVWVZaWUNyQXZ1dmRJe"

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(MySigningKey), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return -1
}
func RequireTokenAuthentication(tokenstring string) (bool, int) {
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(MySigningKey), nil
	})
	return token.Valid, getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"])
}
