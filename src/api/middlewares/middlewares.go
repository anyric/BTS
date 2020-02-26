package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

//Exception custom message
type Exception struct {
	Message string `json:"message"`
}

//SetMiddleWareLogger log user requests
func SetMiddleWareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("")
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

//SetMiddleWareJSON to return requests responses
func SetMiddleWareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//ValidateMiddleWare checks token authorization
func ValidateMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		w.Header().Set("Content-Type", "application/json")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			claims := jwt.MapClaims{}
			if len(bearerToken) == 2 {
				token, error := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(jwt.SigningMethod); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(os.Getenv("TOKEN_SALT")), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					next(w, r)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			
		}
	}
}
