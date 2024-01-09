package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controller "github.com/AkbarFikri/signconnect_backend/Controller"
)

func UserRoutes(route *gin.RouterGroup) {
	route.GET("/profile", controller.GetUserProfile)
}

func AuthRoutes(route *gin.RouterGroup) {
	route.GET("/auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Gooddd"})
	})
	route.POST("/signup", controller.Signup)
	route.POST("/signin", controller.Signin)
}

func LeaderboardRoutes(route *gin.RouterGroup) {
	route.GET("/", controller.GetLeaderboard)
	route.GET("/:limit", controller.GetLeaderboard)
}

func GamesRoutes(route *gin.RouterGroup) {
	route.GET("/questions/:level", controller.GetQuestionsByLevel)
	route.POST("/answer/:level", controller.PostAnswerByLevel)
	route.POST("/tambahsoal", controller.AddQuestions)
}

func HomeRoutes(route *gin.Engine) {
	route.GET("/", controller.HomeAPI)
}

func DictionaryRoutes(route *gin.RouterGroup) {
	route.GET("/kategori", controller.GetDictionaryKategori)
	route.GET("/:kategori/list", controller.GetDictionaryListByKategori)
	route.GET("/:kategori/:id", controller.GetDictionaryValueByKategoriAndID)
}

func LembagaRoutes(route *gin.RouterGroup) {
	route.GET("/", controller.GetDaftarLembaga)
	route.GET("/:id", controller.GetDetailLembaga)
	route.POST("/:id", controller.SendVolunteerApplication)
}
