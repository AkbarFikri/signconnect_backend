package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controller "github.com/AkbarFikri/signconnect_backend/Controller"

)

func UserRoutes(route *gin.RouterGroup) {
	route.GET("/anjay", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gooddd"})
	})
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gooddd"})
	})
}

func AuthRoutes(route *gin.RouterGroup) {
	route.GET("testing", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gooddd"})
	})
	route.POST("/signup", controller.Signup)
	route.POST("/signin", controller.Signin)
}

func LeaderboardRoutes(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gooddd"})
	})
}

func HomeRoutes(route *gin.Engine) {
	route.GET("/", controller.HomeAPI)
}
