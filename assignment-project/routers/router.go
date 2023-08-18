package routers

import (
	"assignment-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetOrders)
	orderRouter := router.Group("/order")
	{
		orderRouter.POST("/", controllers.CreateOrder)
		orderRouter.PUT("/:orderId", controllers.UpdateOrder)
		orderRouter.DELETE("/:orderId", controllers.DeleteOrder)
	}
	return router
}
