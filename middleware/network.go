package middleware

import (
	"github.com/gin-gonic/gin"
)

// NetworkMiddleware : switch network between testnet and mainnet
func NetworkMiddleware(c *gin.Context) {
	network := c.GetHeader("Network")
	c.Set("network", network)
	c.Next()
}
