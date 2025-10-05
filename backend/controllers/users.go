package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/auth"
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
		c.JSON(500, gin.H{"error": "Error to connect to database"})
	}

	result := db.Find(&users)
	c.JSON(200, gin.H{"users": result})
}
func GetUserByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get user by ID"})
}
func CreateUser(c *gin.Context) {
	dbconf := config.LoadDatabase()

	db, err := database.ConnectDB(dbconf)
	if err != nil {
		log.Print(err)
		log.Print("Error to connect to database")
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "customer"

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	log.Printf("User created: %v", user)
	c.JSON(201, "User created successfully")
}

func Login(c *gin.Context) {
	var err error

	user := models.User{}

	dbconf := config.LoadDatabase()

	db, err := database.ConnectDB(dbconf)
	if err != nil {
		log.Print(err)
		log.Print("Error to connect to database")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	log.Printf("User logged in: %v", user)
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token"})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update user by ID"})
}
func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete user by ID"})
}
