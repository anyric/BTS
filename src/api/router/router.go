package router

import (
	"github.com/gorilla/mux"
	"github.com/anyric/bts/src/api/router/routes"
)

//New creates request routers
func New()  *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}