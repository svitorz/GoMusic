package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/auth"
	"github.com/svitorz/GoMusic/backend/models"
	"github.com/svitorz/GoMusic/backend/repository"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	result := repository.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

func GetUserByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get user by ID"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "customer"

	// Encrypt the password before saving
	encryptedPassword, err := encryptPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error encrypting password"})
		return
	}
	user.Password = encryptedPassword

	// Save the user to the database
	result := repository.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	log.Printf("User created: %v", user)
	c.JSON(201, "User created successfully")
}

func Login(c *gin.Context) {
	var err error

	// Estrutura separada para input
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	result := repository.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Aqui, user.Password é o hash do banco; input.Password é a senha digitada
	if passwordIsValid, err := verifyPassword(input.Password, user.Password); err != nil || !passwordIsValid {
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
