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

func CreateStudent(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	Student := models.Student{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Student)
	} else {
		ctx.ShouldBind(&Student)
	}

	err := db.Debug().Create(&Student).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Student,
	})
}

func GetStudents(ctx *gin.Context) {
	db := database.GetDB()
	data := []models.Student{}

	err := db.Debug().Preload("Scores").Find(&data).Error
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

func UpdateStudent(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	studentId, _ := strconv.Atoi(ctx.Param("studentId"))
	studentUpdate := models.Student{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&studentUpdate)
	} else {
		ctx.ShouldBind(&studentUpdate)
	}

	//proses mencari data student sesuai id
	var existingStudent models.Student
	if err := db.Preload("Scores").Where("id = ?", studentId).First(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Student data not found",
		})
		return
	}

	tx := db.Begin()
	for i, newScore := range studentUpdate.Scores {
		if i < len(existingStudent.Scores) {
			//menginisiasi ulang data score pada existing student
			existingStudent.Scores[i].AssignmentTitle = newScore.AssignmentTitle
			existingStudent.Scores[i].Description = newScore.Description
			existingStudent.Scores[i].Score = newScore.Score

			//menyimpan data score yang baru
			err := tx.Save(&existingStudent.Scores[i]).Error
			if err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   err.Error(),
					"message": "Failed to update data score",
				})
				return
			}
		} else {
			existingStudent.Scores = append(existingStudent.Scores, newScore)
			err := tx.Create(&existingStudent.Scores[i]).Error
			if err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   err.Error(),
					"message": "Failed to create new scores",
				})
				return
			}
		}
	}

	//proses mengupdate data yang sudah ada dengan data yang baru
	err := tx.Model(&existingStudent).Updates(&studentUpdate).Error
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"message": "Failed to update data student",
		})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"Data": existingStudent,
	})

}

func DeleteStudent(ctx *gin.Context) {
	db := database.GetDB()
	studentId := ctx.Param("studentId")
	Student := models.Student{}

	studentIdDelete, _ := strconv.Atoi(studentId)

	check := db.First(&Student, studentIdDelete).Error
	if check != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errro":   check.Error(),
			"message": "Student ID is not valid",
		})
		return
	}

	err := db.Select(clause.Associations).Delete(&Student).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data has been deleted",
	})

}
