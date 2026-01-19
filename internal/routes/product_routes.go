package routes

import (
	"github.com/Simok666/ecommerce-app.git/internal/controllers"
	"github.com/Simok666/ecommerce-app.git/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProduct)
	router.GET("/products/:id", controllers.GetProductByID)

	admin := router.Group("/admin/products")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.AdminOnly(),
	)
	{
		admin.POST("", controllers.CreateProduct)
		admin.PUT("/:id", controllers.UpdateProduct)
		admin.DELETE("/:id", controllers.DeleteProduct)
		admin.DELETE("/images/:id", controllers.DeleteProductImage)
	}
}
