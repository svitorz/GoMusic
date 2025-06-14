package router

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	version := router.Group("/v1")
	getArtistasRoutes(version)
}
