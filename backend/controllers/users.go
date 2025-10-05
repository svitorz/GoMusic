package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/config"
	"github.com/svitorz/GoMusic/backend/database"
	"github.com/svitorz/GoMusic/backend/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	dbconf := config.LoadDatabase()

	db, err := database.ConnectDB(dbconf)
	if err != nil {
		log.Print(err)
		log.Print("Error to connect to database")
	}

	result := db.Find(&users)
	c.JSON(200, gin.H{"message": "Success!", "users": result})
}
func GetUserByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get user by ID"})
}
func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create a new user"})
}
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update user by ID"})
}
func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete user by ID"})
}
