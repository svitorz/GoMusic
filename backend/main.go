package main

import (
	"fmt"
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

	db, err := database.ConnectDb(dbconf)
	if err != nil {
		log.Print(err)
		log.Print("Error to connect to database")
	}

	fmt.Println(db)
	appRoutes.Run()
}
