package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.Request.Header.Get("Authorization")
		if val != "secret-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unutorized"})
		}
		c.Next()
	}
}
