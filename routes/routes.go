package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "outblock.io/go-server/demo/controllers"
	"outblock.io/go-server/demo/middleware"
)

func Routes(router *gin.Engine) {
	wallet := new(controllers.WalletController)

	// configure firebase to put in the ctx of requests
	privateV1 := router.Group("/v1")
	// set firebase auth to gin context with a middleware to all incoming request
	privateV1.Use(middleware.Authtoken)
	privateV1.Use(middleware.NetworkMiddleware)
	{
		privateV1.POST("/address", wallet.CreateAddress)
		privateV1.GET("/address", wallet.Getrecord)
		privateV1.POST("/address/testnet", wallet.CreateAddressTest)
		privateV1.GET("/address/testnet", wallet.GetrecordTest)

	}

	router.GET("/", welcome)

	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
