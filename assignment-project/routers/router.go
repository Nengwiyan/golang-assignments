package routers

import (
	"assignment-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	orderRouter := router.Group("/orders")
	{
		orderRouter.POST("/", controllers.CreateOrder)
		orderRouter.GET("/", controllers.GetOrders)
		orderRouter.PUT("/:orderId", controllers.UpdateOrder)
		orderRouter.DELETE("/:orderId", controllers.DeleteOrder)
	}
	return router
}
