package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route describe the data structure for the routes
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Load slices of the route
func Load() []Route {
	routes := usersRoutes
	return routes
}

// SetupRoutes using the gorilla mux router
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
