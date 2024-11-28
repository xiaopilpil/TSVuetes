package router

import (
	"TSVuetesApp/controllors"
	"TSVuetesApp/middlewares"

	// "TSVuetesApp/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     config.AppConfig.Cors.AllowOrigins,
	// 	AllowMethods:     config.AppConfig.Cors.AllowMethods,
	// 	AllowHeaders:     config.AppConfig.Cors.AllowHeaders,
	// 	ExposeHeaders:    config.AppConfig.Cors.ExposeHeaders,
	// 	AllowCredentials: config.AppConfig.Cors.AllowCredentials,
	// 	MaxAge:           config.AppConfig.Cors.MaxAge,
	// }))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://47.242.39.174:3000"},
		AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllors.Login)
		auth.POST("/register", controllors.Register)
	}

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/classroom", controllors.CreateClassroom)
		api.GET("/classroom", controllors.Getclassroom)
		api.POST("/course", controllors.CreateCourse)
		api.GET("/course", controllors.GetCourse)
		api.PUT("/course", controllors.UpdateCourseState)
		api.GET("/course/:id", controllors.GetCourseByID)
		api.PUT("/ps", controllors.UpdateCoursePs)
		api.POST("/userinf", controllors.Getinformation)
	}
	return r
}
