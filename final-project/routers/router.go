package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	adminRouter := router.Group("/auth")
	{
		adminRouter.POST("/register", controllers.CreateAdmin)
		adminRouter.POST("/login", controllers.AdminLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:uuid", controllers.GetByUUID)
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:uuid", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:uuid", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	variantRouter := router.Group("/products/variants")
	{
		variantRouter.GET("/", controllers.GetAllVariant)
		variantRouter.GET("/:uuid", controllers.GetVariantByUUID)
		variantRouter.Use(middlewares.Authentication())
		variantRouter.POST("/", middlewares.VariantAuthorizationPost(), controllers.CreateVariant)
		variantRouter.PUT("/:uuid", middlewares.VariantAuthorization(), controllers.UpdateVariant)
		variantRouter.DELETE("/:uuid", middlewares.VariantAuthorization(), controllers.DeleteVariant)
	}
	return router
}
