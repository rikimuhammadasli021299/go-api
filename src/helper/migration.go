package helper

import (
	"go-api/src/config"
	"go-api/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Product{})
}