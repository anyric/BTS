package routes

import (
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route describe the data structure for the routes
type Route struct {
	URI          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Load slices of the route
func Load() []Route {
	routes := usersRoutes
	routes = append(routes, authRoutes...) 
	return routes
}

// SetupRoutes using the gorilla mux router
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}

// SetupRoutesWithMiddlewares for logging user requests
func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(route.URI, middlewares.ValidateMiddleWare(
				middlewares.SetMiddleWareLogger(
					middlewares.SetMiddleWareJSON(route.Handler)),
			)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.SetMiddleWareLogger(
				middlewares.SetMiddleWareJSON(route.Handler)),
			).Methods(route.Method)
		}
	}
	return r
}
