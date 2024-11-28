package controllors

import (
	"TSVuetesApp/global"
	"TSVuetesApp/models"
	"TSVuetesApp/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var teacher models.Teacher
	//绑定teacher的JSON格式
	if err := ctx.ShouldBindJSON(&teacher); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//生成teacher的哈西密码
	hashedPwd, err := utils.HashPassword(teacher.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	teacher.Password = hashedPwd
	//对id生成一个密钥
	token, err := utils.GenerateJWT(teacher.Teacher_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// // 数据库创建表
	// if err := global.Db.AutoMigrate(&teacher); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// 创建记录
	if err := global.Db.Create(&teacher).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(ctx *gin.Context) {
	var input struct {
		Teacher_id string `json:"teacher_id"`
		Password   string `json:"password"`
	}
	//绑定JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var teacher models.Teacher

	if err := global.Db.Where("teacher_id = ?", input.Teacher_id).First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "wrong credentials"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	// 验证密码
	if !utils.CheckPassword(input.Password, teacher.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}
	token, err := utils.GenerateJWT(input.Teacher_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Getinformation(ctx *gin.Context) {
	// 定义接受 JSON 请求体的所有参数
	var input struct {
		TeacherID string `json:"teacher_id"` // 请求体中的 teacher_id
	}

	// 定义返回的字段
	var teacherInfo struct {
		TeacherID string `json:"teacher_id"`
		Name      string `json:"name"`
		Gender    string `json:"gender"`
		Faculty   string `json:"faculty"`
		Position  int    `json:"position"`
	}

	// 绑定 JSON 请求体
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// 查询数据库，映射到指定字段
	if err := global.Db.Model(&models.Teacher{}).
		Select("teacher_id, name, gender, faculty, position").
		Where("teacher_id = ?", input.TeacherID).
		Scan(&teacherInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher_id: teacher does not exist"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query teacher_id", "details": err.Error()})
		}
		return
	}

	// 返回筛选后的教师信息
	ctx.JSON(http.StatusOK, teacherInfo)
}
