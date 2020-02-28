package router

import (
	"api/router/routes"

	"github.com/gorilla/mux"
)

//New creates request routers
func New()  *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}