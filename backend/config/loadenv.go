package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Name         string
	User         string
	Secret       string
	Host         string
	Port         int
	ExternalPort int
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to handle .env", err)
	}
}

func LoadDatabase() *Database {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432
	}

	externalPort, err := strconv.Atoi(os.Getenv("DB_PORT_EXTERNAL"))
	if err != nil {
		externalPort = 12470
	}

	return &Database{
		Host:         os.Getenv("DB_HOST"),
		User:         os.Getenv("DB_USER"),
		Secret:       os.Getenv("DB_PASSWORD"),
		Name:         os.Getenv("DB_NAME"),
		Port:         port,
		ExternalPort: externalPort,
	}
}
