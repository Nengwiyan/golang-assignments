package routers

import (
	"pertemuan-sembilan/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	productRouter := router.Group("/products")
	{
		productRouter.POST("/", controllers.CreateProduct)
	}

	return router
}
