package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"massage": "SignConnect API Apps Starting",
	})
}
