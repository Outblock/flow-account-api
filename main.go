package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "outblock.io/go-server/demo/docs"
	"outblock.io/go-server/demo/middleware"
	models "outblock.io/go-server/demo/models"
	routes "outblock.io/go-server/demo/routes"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect DB
	models.Connect()
	// Init Router
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", middleware.AuthSwaggo, ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
