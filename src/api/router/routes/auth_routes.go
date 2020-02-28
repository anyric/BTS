package routes

import (
	"net/http"
	"api/controllers"
)
var authRoutes =[]Route{
	Route{
		URI: "/login",
		Method: http.MethodPost,
		Handler: controllers.Login,
		AuthRequired: false,
	},
}