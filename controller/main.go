package controller

import (
	"http_request/class-management/config"
	dbservice "http_request/class-management/services"
)

type Controller struct {
	DatabaseService dbservice.DatabaseService
	appConfig       config.AppConfig
}

func InitController() Controller {
	appConfig := config.LoadAppConfig()

	return Controller{
		DatabaseService: dbservice.InitDatabase(),
		appConfig:       appConfig,
	}
}
