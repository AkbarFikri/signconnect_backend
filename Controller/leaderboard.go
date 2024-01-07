package controller

import (
	"net/http"
	"strconv"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	models "github.com/AkbarFikri/signconnect_backend/Models"
	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	limitStr := c.Param("limit")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid limit parameter",
			})
			return
		}

		var users []models.User
		database.DB.Order("experience desc").Limit(limit).Find(&users)

		leaderboard := make([]gin.H, len(users))
		for i, user := range users {
			leaderboard[i] = gin.H{
				"id":         user.ID,
				"username":   user.Username,
				"experience": user.Experience,
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"leaderboard": leaderboard,
		})
	} else {
		var users []models.User
		database.DB.Order("experience desc").Find(&users)

		leaderboard := make([]gin.H, len(users))
		for i, user := range users {
			leaderboard[i] = gin.H{
				"id":         user.ID,
				"username":   user.Username,
				"experience": user.Experience,
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"leaderboard": leaderboard,
		})
	}
}
