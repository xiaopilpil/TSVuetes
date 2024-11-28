package config

import (
	"TSVuetesApp/global"
	"TSVuetesApp/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}
	global.Db = db

	// 数据库创建表
	if err := global.Db.AutoMigrate(&models.Teacher{}); err != nil {
		log.Fatalf("Failed to initialize table, got error: %v", err)
	}

	if err := global.Db.AutoMigrate(&models.Classroom{}); err != nil {
		log.Fatalf("Failed to initialize table, got error: %v", err)
	}

	if err := global.Db.AutoMigrate(&models.Course{}); err != nil {
		log.Fatalf("Failed to initialize table, got error: %v", err)
	}

}
