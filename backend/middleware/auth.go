package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svitorz/GoMusic/backend/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
