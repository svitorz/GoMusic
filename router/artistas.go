package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArtistasRoutes(rg *gin.RouterGroup) {
	artistas := rg.Group("/artistas")

	// TODO: criar funcoes no controller para as rotas
	artistas.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Olá, sua api de músicas está funcionando")
	})

	// artistas.GET(":nickname")
}
