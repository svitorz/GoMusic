package main

import (
	"github.com/svitorz/GoMusic/backend/router"
)

func main() {
	appRoutes := router.SetupRouter()

	appRoutes.Run()
}
