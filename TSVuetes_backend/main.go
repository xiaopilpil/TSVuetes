package main

import (
	"TSVuetesApp/config"
	"TSVuetesApp/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()
	r.Run(":" + config.AppConfig.App.Port)
}
