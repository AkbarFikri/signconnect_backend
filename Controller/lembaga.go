package controller

import (
	"net/http"
	"strconv"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	"github.com/AkbarFikri/signconnect_backend/Models"
	"github.com/gin-gonic/gin"
)

func GetDaftarLembaga(c *gin.Context) {
	var lembagas []models.Lembaga
	database.DB.Find(&lembagas)

	result := make([]gin.H, len(lembagas))
	for i, lembaga := range lembagas {
		result[i] = gin.H{
			"id":                lembaga.Id,
			"nama":              lembaga.Nama,
			"lokasi":            lembaga.Tempat,
			"lembaga_image_url": lembaga.Image_url,
			"minimal_tahun":     lembaga.Min_pengalaman,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetDetailLembaga(c *gin.Context) {
	idStr := c.Param("id")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id parameter",
			})
			return
		}

		var lembaga models.Lembaga
		database.DB.First(&lembaga, id)

		if lembaga.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Lembaga not found",
			})
			return
		}

		result := gin.H{
			"id":                  lembaga.Id,
			"nama":                lembaga.Nama,
			"lokasi":              lembaga.Tempat,
			"lembaga_image_url":   lembaga.Image_url,
			"minimal_tahun":       lembaga.Min_pengalaman,
			"tentang_lembaga":     lembaga.Deskripsi,
			"deskripsi_pekerjaan": lembaga.Pekerjaan,
			"persyaratan":         lembaga.Persyaratan,
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id parameter",
		})
	}
}

func SendVolunteerApplication(c *gin.Context) {
	idStr := c.Param("id")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid lembaga ID parameter",
			})
			return
		}

		var lembaga models.Lembaga
		result := database.DB.First(&lembaga, id)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Lembaga not found",
			})
			return
		}

		userObj, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User information not found",
			})
			return
		}

		userD, ok := userObj.(models.User)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user information",
			})
			return
		}

		userID := userD.ID
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User information not found",
			})
			return
		}

		var user models.User
		userResult := database.DB.First(&user, userID)
		if userResult.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User not found",
			})
			return
		}

		var requestBody struct {
			CvURL string `json:"cv_url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		application := models.Application{
			CvURL:     requestBody.CvURL,
			Username:  user.Username,
			Email:     user.Email,
			LembagaID: lembaga.Id,
			UserID:    int(userID),
		}

		database.DB.Create(&application)

		c.JSON(http.StatusOK, gin.H{
			"result": gin.H{
				"lembaga_id": lembaga.Id,
				"message":    "Success Mengirimkan Lamaran",
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid lembaga ID parameter",
		})
	}
}
