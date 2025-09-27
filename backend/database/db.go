package database

import (
	"fmt"

	"github.com/svitorz/GoMusic/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(database *config.Database) (*gorm.DB, error) {
	var dsn string
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		database.Host,
		database.User,
		database.Secret,
		database.Name,
		// database.Port,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error to connect to postgres")
		return nil, err
	}

	return db, nil
}
