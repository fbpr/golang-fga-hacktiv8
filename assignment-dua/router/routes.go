package router

import (
	"assignment-dua/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	orderRouter := router.Group("/orders")
	{
		orderRouter.POST("/", controllers.CreateOrder)
		orderRouter.GET("/", controllers.GetOrder)
		orderRouter.PUT("/:orderId", controllers.UpdateOrderById)
		orderRouter.DELETE("/:orderId", controllers.DeleteOrderById)
	}

	return router
}