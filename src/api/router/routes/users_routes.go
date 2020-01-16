package routes

import (
	"net/http"

	"github.com/anyric/bts/src/api/controllers"
)
var usersRoutes =[]Route{
	Route{
		URI: "/api/v1/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		URI: "/api/v1/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		URI: "/api/v1/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		URI: "/api/v1/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		URI: "/api/v1/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
	}}