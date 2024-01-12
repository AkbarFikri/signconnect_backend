package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	"github.com/AkbarFikri/signconnect_backend/Models"
)

// GetDictionaryKategori returns a list of dictionary categories.
func GetDictionaryKategori(c *gin.Context) {
	var kategoriList []models.DictionaryKategori
	database.DB.Find(&kategoriList)

	result := make([]gin.H, len(kategoriList))
	for i, kategori := range kategoriList {
		result[i] = gin.H{
			"id":       kategori.Id,
			"kategori": kategori.Kategori,
			"imageURL": kategori.ImageURL,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetDictionaryListByKategori(c *gin.Context) {
	kategoriName := c.Param("kategori")

	// Retrieve category ID using the category name
	var kategori models.DictionaryKategori
	if err := database.DB.Where("kategori = ?", kategoriName).First(&kategori).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	var dictionaryList []models.Dictionary
	database.DB.Where("kategori_id = ?", kategori.Id).Find(&dictionaryList)

	result := make([]gin.H, len(dictionaryList))
	for i, dictionary := range dictionaryList {
		result[i] = gin.H{
			"id":          dictionary.Id,
			"kategori_id": dictionary.KategoriID,
			"huruf":       dictionary.Huruf,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetDictionaryValueByKategoriAndID(c *gin.Context) {
	kategoriName := c.Param("kategori")

	// Retrieve category ID using the category name
	var kategori models.DictionaryKategori
	if err := database.DB.Where("kategori = ?", kategoriName).First(&kategori).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Category not found: %v", err),
		})
		return
	}

	// Parse id parameter as an integer
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is empty",
		})
		return
	}

	valueID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid ID format: %v", err),
		})
		return
	}

	// Retrieve dictionary value using category ID and parsed ID
	var dictionaryValue models.Dictionary
	if err := database.DB.Where("kategori_id = ? AND id = ?", kategori.Id, valueID).First(&dictionaryValue).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Dictionary value not found: %v", err),
		})
		return
	}

	// Return the result
	result := gin.H{
		"id":          dictionaryValue.Id,
		"kategori_id": dictionaryValue.KategoriID,
		"huruf":       dictionaryValue.Huruf,
		"image_url":   dictionaryValue.ImageURL,
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
