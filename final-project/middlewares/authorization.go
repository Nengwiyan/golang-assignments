package middlewares

import (
	"final-project/database"
	models "final-project/models/entity"
	request "final-project/models/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		getUUID := ctx.Param("uuid")

		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))
		Product := models.Product{}

		//Get Admin id of uuid param from product table
		err := db.Select("admin_id").Where("uuid = ?", fmt.Sprintf(`%s`, getUUID)).First(&Product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data not found",
			})
			return
		}

		//compare admin id from table product with admin id from context
		if Product.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func VariantAuthorizationPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		VariantReq := request.VariantRequest{}
		Product := models.Product{}

		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))

		if err := ctx.ShouldBind(&VariantReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		uuidProduct := VariantReq.ProductID
		//get admin id by uuid from table product
		err := db.Select("admin_id").Where("uuid = ?", fmt.Sprintf(`%s`, uuidProduct)).First(&Product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data not found",
			})
			return
		}

		//compare admin id from table product with admin id from context, for authorization
		if Product.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))
		uuidVariant := ctx.Param("uuid")
		Variant := models.Variant{}
		Product := models.Product{}

		//Get id product of uuid variant
		getIdProduct := db.Select("product_id").Where("uuid = ?", fmt.Sprintf(`%s`, uuidVariant)).Find(&Variant).Error
		if getIdProduct != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data ID Product not found",
				"message": getIdProduct.Error(),
			})
			return
		}

		//Get admin Id of id product
		ProductID := Variant.ProductID
		getAdminID := db.Select("admin_id").Where("id = ?", ProductID).First(&Product).Error
		if getAdminID != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Admin ID of Product is not found",
				"message": getAdminID.Error(),
			})
			return
		}

		// compare id product with admin id of context
		if Product.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
