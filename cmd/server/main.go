package main

import (
	"os"

	"github.com/Simok666/ecommerce-app.git/internal/config"
	"github.com/Simok666/ecommerce-app.git/internal/database"
	"github.com/Simok666/ecommerce-app.git/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectDatabase()

	r := gin.Default()

	routes.AuthRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
