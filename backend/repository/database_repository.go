package repository

import (
	"github.com/svitorz/GoMusic/backend/config"
	"github.com/svitorz/GoMusic/backend/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config := config.LoadDatabase()
	var err error
	DB, err = database.ConnectDB(config)
	if err != nil {
		panic("Failed to connect to database!")
	}
}
