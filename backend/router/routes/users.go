package router

import (
	"net/http"

	"github.com/svitorz/GoMusic/backend/controllers"
)

var userRoutes = []Route{
	{
		Path:        "/users",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUsers,
		useAuth:     true,
	},
	{
		Path:        "/users/{id}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserByID,
		useAuth:     true,
	},
	{
		Path:        "/users",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateUser,
		useAuth:     false,
	},
	{
		Path:        "/users/:id",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateUser,
		useAuth:     true,
	},
	{
		Path:        "/users/:id",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUser,
		useAuth:     true,
	},
}
