package controllers

import (
	"net/http"
	"path/filepath"
	"pertemuan-sembilan/database"
	"pertemuan-sembilan/helpers"
	models "pertemuan-sembilan/models/entity"
	request "pertemuan-sembilan/models/request"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	var productReq request.ProductRequest
	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the filename without extension
	fileName := helpers.RemoveExtention(productReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product := models.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		ImageURL:    uploadResult,
	}

	//detect file size
	fileSize := productReq.Image.Size
	maxSize := 5 * 1024 * 1024

	//detect file extention
	fileExt := filepath.Ext(productReq.Image.Filename)
	newExt := strings.ToLower(fileExt)

	//condition to filter file extention and size
	if newExt == ".jpg" || newExt == ".png" || newExt == ".jpeg" {
		if fileSize < int64(maxSize) {
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "File size must be lower than 5MB"})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not upload file. Please insert file with extention .jpg, .png, .jpeg only"})
	}

}
