package router

import (
	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/middleware"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc gin.HandlerFunc
	useAuth     bool
}

func Config(r *gin.Engine) *gin.Engine {
	routes := userRoutes

	for _, route := range routes {

		if route.useAuth {
			r.Handle(route.Method, route.Path, middleware.AuthMiddleware(), route.HandlerFunc)
		} else {
			r.Handle(route.Method, route.Path, route.HandlerFunc)
		}

		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	return r
}
