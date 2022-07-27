package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthSwaggo(c *gin.Context) {
	authCookies, _ := c.Cookie("swagger")
	apiKey := os.Getenv("SWAGGER_KEY")
	if authCookies != apiKey {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Next()
}

func Authtoken(c *gin.Context) {
	apiKey := os.Getenv("API_KEY")
	authorizationToken := c.GetHeader("Authorization")
	if authorizationToken != apiKey {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Next()
}
