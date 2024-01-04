package database

import (
	models "github.com/AkbarFikri/signconnect_backend/Models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
