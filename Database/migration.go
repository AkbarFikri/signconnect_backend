package database

import (
	models "github.com/AkbarFikri/signconnect_backend/Models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Leaderboard{})
	DB.AutoMigrate(&models.Soal{})
	DB.AutoMigrate(&models.Leveling{})
	DB.AutoMigrate(&models.Lembaga{})
	DB.AutoMigrate(&models.UserRole{})
	DB.AutoMigrate(&models.Answer{})
	DB.AutoMigrate(&models.Dictionary{})
	DB.AutoMigrate(&models.DictionaryKategori{})
}
