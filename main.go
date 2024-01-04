package main

import (
	config "github.com/AkbarFikri/signconnect_backend/Config"
	database "github.com/AkbarFikri/signconnect_backend/Database"
	"github.com/AkbarFikri/signconnect_backend/routers"
)

func main() {
	config.SetupConfig()
	database.Database()
	database.Migrate()

	router := routers.SetupRoute()
	router.Run()
}
