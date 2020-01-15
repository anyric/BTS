package router

import (
	"github.com/gorilla/mux"
	"github.com/anyric/bts/src/api/router/routes"
)

func New()  *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}