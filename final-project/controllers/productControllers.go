package controllers

import (
	"final-project/database"
	"final-project/helpers"
	models "final-project/models/entity"
	request "final-project/models/request"
	"fmt"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	ProductReq := request.ProductRequest{}
	adminID := uint(adminData["id"].(float64))

	if err := ctx.ShouldBind(&ProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fileName := helpers.RemoveExtention(ProductReq.Image.Filename)
	uploadResult, err := helpers.UploadFile(ProductReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.New()
	Product := models.Product{
		UUID:     newUUID.String(),
		AdminID:  adminID,
		Name:     ProductReq.Name,
		ImageURL: uploadResult,
	}

	fileExt := filepath.Ext(ProductReq.Image.Filename)
	newExt := strings.ToLower(fileExt)

	if newExt == ".jpg" || newExt == ".png" || newExt == ".jpeg" {
		err = db.Debug().Create(&Product).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": Product,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not upload file. only file with JPG, PNG, and JPEG extention only allowed"})
	}
}

func GetAllProduct(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	search := ctx.Query("search")
	db := database.GetDB()
	Product := []models.Product{}

	//preload variant and admin
	Query := db.Model(&Product).Preload("Variants").Preload("Admin")
	var totalProduct int64
	if search != "" {
		Query = Query.Where("name LIKE ?", "%"+search+"%")
	}

	Query.Count(&totalProduct)
	err := Query.Offset(offset).Limit(limit).Find(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get datas",
			"message": err.Error(),
		})
		return
	}

	page := int(math.Ceil(float64(totalProduct) / float64(limit)))
	if page < 1 {
		page = 1
	}
	lastPage := (offset/limit + 1)

	pagination := models.Pagination{
		LastPage: lastPage,
		Limit:    limit,
		Offset:   int(offset),
		Page:     page,
		Total:    totalProduct,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       Product,
		"pagination": pagination,
	})
}

func GetByUUID(ctx *gin.Context) {
	db := database.GetDB()
	data := models.Product{}
	getUUID := ctx.Param("uuid")

	err := db.Debug().Preload("Variants").Where("uuid = ?", getUUID).Find(&data).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})

}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	ProductReq := request.ProductRequest{}
	adminID := uint(adminData["id"].(float64))
	uuidProduct := ctx.Param("uuid")

	existingProduct := models.Product{}
	err := db.Debug().Where("uuid = ?", uuidProduct).Find(&existingProduct).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Existing Product not found",
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBind(&ProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if ProductReq.Image != nil {
		//remove extention and upload to cloudinary
		fileName := helpers.RemoveExtention(ProductReq.Image.Filename)
		uploadResult, err := helpers.UploadFile(ProductReq.Image, fileName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		existingProduct.ImageURL = uploadResult
	}

	existingProduct.Name = ProductReq.Name
	existingProduct.AdminID = adminID
	err = db.Save(&existingProduct).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": existingProduct,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	getUUID := ctx.Param("uuid")
	Product := models.Product{}

	check := db.First(&Product).Where("uuid = ?", fmt.Sprintf("`%s`", getUUID)).Error
	if check != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": check.Error(),
		})
		return
	}

	err := db.Delete(&Product).Where("uuid = ?", fmt.Sprintf("`%s`", getUUID)).Error
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
