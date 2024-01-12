package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	models "github.com/AkbarFikri/signconnect_backend/Models"

)


func Signup(c *gin.Context) {
	// Get the Email and Password user from req body
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash Password user
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Hash Password",
		})
		return
	}

	//Create The User
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hashed)}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create User",
		})
		return
	}

	leveling := models.Leveling{
		Level: 1,
		UserID: int(user.ID),
		Status: "Ongoing",
	}
	database.DB.Create(&leveling)

	leaderboard := models.Leaderboard{
		UserID:     int(user.ID),
		Experience: 0,
	}
	database.DB.Create(&leaderboard)

	// Respond send
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      user.ID,
		"massage": "Success Create New User",
	})

}

func Signin(c *gin.Context) {
	// Get the Email and Password user from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// check Email request
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// compare password sent with password in database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Token",
		})
		return
	}

	// set cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("X-Authorization", tokenString, 3600*3, "", "", false, true)

	// return token
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"token": tokenString,
	})
}
