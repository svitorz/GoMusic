package router

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	artistas := router.Group("/v1")
	getArtistasRoutes(artistas)
}
