package controllers

import (
	"final-project/database"
	models "final-project/models/entity"
	request "final-project/models/request"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))

	VariantReq := request.VariantRequest{}
	Product := models.Product{}
	Product.AdminID = adminID

	if err := ctx.ShouldBind(&VariantReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	uuidproduct := VariantReq.ProductID
	//Get id product by uuid, to assing for variant.productid
	getIDProduct := db.Select("id").Where("uuid = ?", fmt.Sprintf(`%s`, uuidproduct)).First(&Product).Error
	if getIDProduct != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Data ID Product not found",
			"message": getIDProduct.Error(),
		})
		return
	}

	newUUID := uuid.New()
	Variant := models.Variant{
		UUID:        newUUID.String(),
		VariantName: VariantReq.VariantName,
		Quantity:    VariantReq.Quantity,
		ProductID:   Product.ID,
	}

	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create variant",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Variant,
	})
}

func GetAllVariant(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	search := ctx.Query("search")
	db := database.GetDB()
	Variant := []models.Variant{}

	Query := db.Model(&Variant)
	var totalVariant int64
	if search != "" {
		Query = Query.Where("variant_name LIKE ?", "%"+search+"%")
	}

	Query.Count(&totalVariant)
	err := Query.Offset(offset).Limit(limit).Find(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not show data variant",
			"message": err.Error(),
		})
		return
	}

	page := int(math.Ceil(float64(totalVariant) / float64(limit)))
	if page < 1 {
		page = 1
	}
	lastPage := (offset/limit + 1)

	pagination := models.Pagination{
		LastPage: lastPage,
		Limit:    limit,
		Offset:   int(offset),
		Page:     page,
		Total:    totalVariant,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       Variant,
		"pagination": pagination,
	})
}

func GetVariantByUUID(ctx *gin.Context) {
	db := database.GetDB()
	Data := models.Variant{}
	getUUID := ctx.Param("uuid")

	err := db.Debug().Where("uuid = ?", getUUID).Find(&Data).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Data,
	})
}

func UpdateVariant(ctx *gin.Context) {
	db := database.GetDB()
	VariantReq := request.VariantRequest{}
	uuidVariant := ctx.Param("uuid")

	existingVariant := models.Variant{}
	err := db.Debug().Where("uuid = ?", uuidVariant).Find(&existingVariant).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Existing Product not found",
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBind(&VariantReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	existingVariant.VariantName = VariantReq.VariantName
	existingVariant.Quantity = VariantReq.Quantity
	err = db.Save(&existingVariant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": existingVariant,
	})

}

func DeleteVariant(ctx *gin.Context) {
	db := database.GetDB()
	getUUID := ctx.Param("uuid")
	Variant := models.Variant{}

	check := db.Debug().Where("uuid = ?", fmt.Sprintf(`%s`, getUUID)).Find(&Variant).Error
	if check != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": check.Error(),
		})
		return
	}

	err := db.Delete(&Variant).Where("uuid = ?", fmt.Sprintf("`%s`", getUUID)).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
	})
}
