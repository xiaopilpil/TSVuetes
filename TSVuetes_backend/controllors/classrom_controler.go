package controllors

import (
	"TSVuetesApp/global"
	"TSVuetesApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClassroom(ctx *gin.Context) {
	var classroom models.Classroom

	if err := ctx.ShouldBindJSON(&classroom); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if err := global.Db.AutoMigrate(&classroom); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := global.Db.Create(&classroom).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, classroom)
}

func Getclassroom(ctx *gin.Context) {
	var classrooms []models.Classroom

	if err := global.Db.Find(&classrooms).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, classrooms)
}
