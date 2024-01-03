package connections

import models "github.com/AkbarFikri/signconnect_backend/Models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
