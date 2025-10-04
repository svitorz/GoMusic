package database

import (
	"fmt"

	"github.com/svitorz/GoMusic/backend/config"
	"github.com/svitorz/GoMusic/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(database *config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		database.Host,
		database.User,
		database.Secret,
		database.Name,
		// database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database successfully")

	fmt.Println("running migrations...")
	db.Migrator().DropTable(&models.User{}, &models.Album{}, &models.Music{}, &models.Playlist{})
	db.AutoMigrate(&models.User{}, &models.Album{}, &models.Music{}, &models.Playlist{})

	return db, nil
}
