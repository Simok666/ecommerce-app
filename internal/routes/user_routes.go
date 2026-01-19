package routes

import (
	"github.com/Simok666/ecommerce-app.git/internal/controllers"
	"github.com/Simok666/ecommerce-app.git/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")

	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", controllers.Profile)
	}
}
