package router

import (
	"github.com/gin-gonic/gin"
	routes "github.com/svitorz/GoMusic/backend/router/routes"
)

func SetupRouter() *gin.Engine {
	return routes.Config(gin.Default())
}
