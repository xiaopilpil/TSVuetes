package controllors

import (
	"TSVuetesApp/global"
	"TSVuetesApp/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCourse(ctx *gin.Context) {
	var course models.Course

	// 绑定JSON数据
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// 设置初始状态和同意 ID
	course.State = 0
	course.Agree_id = ""
	course.Ps = ""
	// if err := global.Db.AutoMigrate(&course); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 检查 teacher_id 是否存在
	if err := global.Db.First(&models.Teacher{}, "teacher_id = ?", course.Teacher_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher_id: teacher does not exist"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check teacher_id", "details": err.Error()})
		}
		return
	}

	// 检查 room_id 是否存在
	if err := global.Db.First(&models.Classroom{}, "room_id = ?", course.Room_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room_id: room does not exist"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check room_id", "details": err.Error()})
		}
		return
	}

	// 插入数据
	if err := global.Db.Create(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Course with the same TeacherID and RoomID already exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course", "details": err.Error()})
		}
		return
	}

	// 成功响应
	ctx.JSON(http.StatusOK, gin.H{"message": "Course created successfully"})
}

func GetCourse(ctx *gin.Context) {
	var course []models.Course

	if err := global.Db.Find(&course).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func UpdateCourseState(ctx *gin.Context) {
	// 定义一个结构体来接收 JSON 请求体中的所有参数
	var requestBody struct {
		TeacherID string `json:"teacher_id"` // 从 JSON 中获取 teacher_id
		RoomID    string `json:"room_id"`    // 从 JSON 中获取 room_id
		State     int    `json:"state"`      // 从 JSON 中获取 state
	}

	// 绑定 JSON 请求体
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// 检查是否提供了必须的参数
	if requestBody.TeacherID == "" || requestBody.RoomID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "teacher_id and room_id are required"})
		return
	}

	// 首先检查是否存在指定的 teacher_id 和 room_id 的记录
	var course models.Course
	if err := global.Db.Where("teacher_id = ? AND room_id = ?", requestBody.TeacherID, requestBody.RoomID).First(&course).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果记录不存在，返回 404 错误
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found with specified teacher_id and room_id"})
		} else {
			// 其他错误（例如数据库连接问题），返回 500 错误
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query course", "details": err.Error()})
		}
		return
	}

	// 如果记录存在，执行更新操作
	if err := global.Db.Model(&course).
		Update("state", requestBody.State).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update state", "details": err.Error()})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Course state updated successfully",
		"teacher_id": requestBody.TeacherID,
		"room_id":    requestBody.RoomID,
		"new_state":  requestBody.State,
	})
}

func GetCourseByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var course []models.Course

	if err := global.Db.Where("teacher_id = ?", id).Find(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func UpdateCoursePs(ctx *gin.Context) {
	// 定义一个结构体来接收 JSON 请求体中的所有参数
	var requestBody struct {
		TeacherID string `json:"teacher_id"` // 从 JSON 中获取 teacher_id
		RoomID    string `json:"room_id"`    // 从 JSON 中获取 room_id
		Ps        string `json:"ps"`         // 从 JSON 中获取 ps
	}

	// 绑定 JSON 请求体
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// 检查是否提供了必须的参数
	if requestBody.TeacherID == "" || requestBody.RoomID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "teacher_id and room_id are required"})
		return
	}

	// 首先检查是否存在指定的 teacher_id 和 room_id 的记录
	var course models.Course
	if err := global.Db.Where("teacher_id = ? AND room_id = ?", requestBody.TeacherID, requestBody.RoomID).First(&course).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果记录不存在，返回 404 错误
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found with specified teacher_id and room_id"})
		} else {
			// 其他错误（例如数据库连接问题），返回 500 错误
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query course", "details": err.Error()})
		}
		return
	}

	// 如果记录存在，执行更新操作
	if err := global.Db.Model(&course).
		Update("ps", requestBody.Ps).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update state", "details": err.Error()})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Course state updated successfully",
		"teacher_id": requestBody.TeacherID,
		"room_id":    requestBody.RoomID,
		"ps":         requestBody.Ps,
	})
}
