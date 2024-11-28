package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn string
	}
	Cors struct {
		AllowOrigins     []string
		AllowMethods     []string
		AllowHeaders     []string
		ExposeHeaders    []string
		AllowCredentials bool
		MaxAge           time.Duration
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	initDB()
}
