package controllers

import (
	"assignment-project/database"
	"assignment-project/helpers"
	"assignment-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

var (
	appJSON = "application/json"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	Order := models.Order{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Order)
	} else {
		ctx.ShouldBind(&Order)
	}

	err := db.Debug().Create(&Order).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Order,
	})
}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()
	data := []models.Order{}

	err := db.Debug().Preload("Items").Find(&data).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	OrderUpdate := models.Order{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&OrderUpdate)
	} else {
		ctx.ShouldBind(&OrderUpdate)
	}

	//proses mencari data order sesuai id
	var existingOrder models.Order
	if err := db.Preload("Items").Where("id = ?", orderId).First(&existingOrder).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Order data not found",
		})
		return
	}

	tx := db.Begin()
	for i, newItem := range OrderUpdate.Items {
		var existingItem models.Item
		//proses mencari data item berdasarkan order_id
		err := tx.Where("order_id = ?", orderId).First(&existingItem).Error
		if err == nil {
			//inisiasi dengan data baru
			existingItem.Name = newItem.Name
			existingItem.Description = newItem.Description
			existingItem.Quantity = newItem.Quantity

			existingOrder.Items[i].Name = newItem.Name
			existingOrder.Items[i].Description = newItem.Description
			existingOrder.Items[i].Quantity = newItem.Quantity

			//menyimpan data yang baru
			err := tx.Save(&existingItem).Error
			if err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   err.Error(),
					"message": "Failed to update data item",
				})
				return
			}
		}
	}

	//proses mengupdate data yang sudah ada dengan data yang baru
	err := tx.Model(&existingOrder).Updates(&OrderUpdate).Error
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"message": "Failed to update data order",
		})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"Data": existingOrder,
	})

}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderId := ctx.Param("orderId")
	Order := models.Order{}

	orderIdDelete, _ := strconv.Atoi(orderId)

	check := db.First(&Order, orderIdDelete).Error
	if check != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errro":   check.Error(),
			"message": "Order ID is not valid",
		})
		return
	}

	err := db.Select(clause.Associations).Delete(&Order).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data has been deleted",
	})

}
