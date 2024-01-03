package main

import (
	"github.com/gin-gonic/gin"

	connections "github.com/AkbarFikri/signconnect_backend/Connections"
	controller "github.com/AkbarFikri/signconnect_backend/Controller"
	middleware "github.com/AkbarFikri/signconnect_backend/Middleware"

)

func init() {
	connections.LoadEnv()
	connections.Database()
	connections.MigrateDB()
}

func main() {
	r := gin.Default()
	r.GET("/", middleware.AuthAPIKey, controller.HomeAPI)
	r.POST("/auth/signup", middleware.AuthAPIKey, controller.Signup)
	r.POST("/auth/signin", middleware.AuthAPIKey, controller.Signin)
	r.Run()
}
