package routers

import (
	"assignment-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.GET("/students", controllers.GetStudents)
	orderRouter := router.Group("/student")
	{
		orderRouter.POST("/", controllers.CreateStudent)
		orderRouter.PUT("/:studentId", controllers.UpdateStudent)
		orderRouter.DELETE("/:studentId", controllers.DeleteStudent)
	}
	return router
}
