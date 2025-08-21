package main

import (
	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/router"
)

func main() {
	appRoutes := gin.Default()

	router.SetupRouter(appRoutes)

	appRoutes.Run(":8080")
}
