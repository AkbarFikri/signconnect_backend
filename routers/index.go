package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/AkbarFikri/signconnect_backend/routers/Middleware"
)

func SetupRoute() *gin.Engine {
	router := gin.New()

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthAPIKey)

	auth := router.Group("/auth")
	AuthRoutes(auth)

	user := router.Group("/user", middleware.AuthJWTToken)
	UserRoutes(user)

	leaderboard := router.Group("/leaderboard", middleware.AuthJWTToken)
	LeaderboardRoutes(leaderboard)

	games := router.Group("/game", middleware.AuthJWTToken)
	GamesRoutes(games)

	dictionary := router.Group("/dictionary", middleware.AuthJWTToken)
	DictionaryRoutes(dictionary)

	lembaga := router.Group("/lembaga", middleware.AuthJWTToken)
	LembagaRoutes(lembaga)

	HomeRoutes(router) //routes register
	return router
}
