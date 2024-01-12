package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	models "github.com/AkbarFikri/signconnect_backend/Models"

)

func GetLeaderboard(c *gin.Context) {
	defaultLimit := 10

	limitStr := c.Param("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	var users []models.Leaderboard
	db := database.DB.Preload("User").Order("experience desc").Limit(limit)

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	leaderboard := make([]gin.H, len(users))
	for i, user := range users {
		leaderboard[i] = gin.H{
			"id":         user.Id,
			"username":   user.User.Username,
			"experience": user.Experience,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"leaderboard": leaderboard,
	})
}
