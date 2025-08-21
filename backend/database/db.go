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
		"host=%s user=%s password=%s dbname=%s port=%o sslmode=disable TimeZone=Asia/Shanghai",
		database.Host,
		database.User,
		database.Secret,
		database.Name,
		database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error to connect to postgres")
		return nil, err
	}

	return db, nil
}
