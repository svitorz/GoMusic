package main

import (
	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/router"
)

func main() {
	appRoutes := gin.Default()

	router.SetupRouter(appRoutes)

	appRoutes.Run()
}
