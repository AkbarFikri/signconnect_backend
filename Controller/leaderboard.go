package controller

import (
	"net/http"
	"strconv"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	models "github.com/AkbarFikri/signconnect_backend/Models"
	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	// Default limit if not provided
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
