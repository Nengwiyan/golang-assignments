package controllers

// import (
// 	"assignment-project/database"
// 	"assignment-project/helpers"
// 	"assignment-project/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func createItem(ctx *gin.Context) {
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(ctx)
// 	Item := models.Item{}

// 	if contentType == appJSON {
// 		ctx.ShouldBindJSON(&Item)
// 	} else {
// 		ctx.ShouldBind(&Item)
// 	}

// 	err := db.Debug().Create(&Item).Error
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": Item,
// 	})
// }
