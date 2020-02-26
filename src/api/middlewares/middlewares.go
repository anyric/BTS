package middlewares

import "net/http"

import "log"

import "fmt"

//SetMiddleWareLogger log user requests
func SetMiddleWareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("")
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

//SetMiddleWareJSON to return requests responses
func SetMiddleWareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}