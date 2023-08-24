package controllers

import (
	"final-project/database"
	"final-project/helpers"
	models "final-project/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAdmin(ctx *gin.Context) {
	db := database.GetDB()
	Admin := models.Admin{}

	if err := ctx.ShouldBind(&Admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newUUID := uuid.New()
	Admin.UUID = newUUID.String()

	err := db.Debug().Create(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": Admin,
	})
}

func AdminLogin(ctx *gin.Context) {
	db := database.GetDB()
	Admin := models.Admin{}
	var password string

	if err := ctx.ShouldBind(&Admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	password = Admin.Password

	err := db.Debug().Where("email = ?", Admin.Email).Take(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(Admin.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(Admin.ID, Admin.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
