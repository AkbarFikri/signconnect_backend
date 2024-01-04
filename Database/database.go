package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/AkbarFikri/signconnect_backend/Config"

)

var (
	DB *gorm.DB
	err error
)

func Database() {
	DSN := config.DbConfiguration()
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Can't Connect to Database")
	}
}
