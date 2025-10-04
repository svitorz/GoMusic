package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/config"
	"github.com/svitorz/GoMusic/backend/database"
	"github.com/svitorz/GoMusic/backend/router"
)

func main() {
	appRoutes := gin.Default()

	router.SetupRouter(appRoutes)

	dbconf := config.LoadDatabase()

	_, err := database.ConnectDB(dbconf)
	if err != nil {
		log.Print(err)
		log.Print("Error to connect to database")
	}

	appRoutes.Run()
}
