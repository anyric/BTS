package routes

import (
	"net/http"
	"api/controllers"
)
var usersRoutes =[]Route{
	Route{
		URI: "/api/v1/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
		AuthRequired: true,
	},
	Route{
		URI: "/api/v1/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
		AuthRequired: false,
	},
	Route{
		URI: "/api/v1/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
		AuthRequired: true,
	},
	Route{
		URI: "/api/v1/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		AuthRequired: true,
	},
	Route{
		URI: "/api/v1/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
		AuthRequired: true,
	}}