package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArtistasRoutes(rg *gin.RouterGroup) {
	artistas := rg.Group("/artistas")

	artistas.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, "Olá, sua api de músicas está funcionando")
	})
}
