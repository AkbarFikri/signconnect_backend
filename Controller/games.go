package controller

import (
	"net/http"
	"strconv"

	database "github.com/AkbarFikri/signconnect_backend/Database"
	"github.com/AkbarFikri/signconnect_backend/Models"
	"github.com/gin-gonic/gin"
)

func GetQuestionsByLevel(c *gin.Context) {
	levelStr := c.Param("level")

	if levelStr != "" {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid level parameter",
			})
			return
		}

		var questions []models.Soal
		database.DB.Where("level = ?", level).Limit(10).Find(&questions)

		result := make([]gin.H, len(questions))
		for i, question := range questions {
			result[i] = gin.H{
				"id":              question.Id,
				"level":           question.Level,
				"questions_image": question.QuestionImg_url,
				"questions":       question.Question,
				"ans1":            question.Answer_1,
				"ans2":            question.Answer_2,
				"ans3":            question.Answer_3,
				"ans4":            question.Answer_4,
				"correct_ans":     question.Correct_answer,
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid level parameter",
		})
	}
}

func PostAnswerByLevel(c *gin.Context) {
    levelStr := c.Param("level")

    if levelStr != "" {
        level, err := strconv.Atoi(levelStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid level parameter",
            })
            return
        }

        // Fetch questions based on the level from the database
        var soals []models.Soal
        database.DB.Where("level = ?", level).Find(&soals)

        var requestBody struct {
            UserAnswer []struct {
                ID         int    `json:"id"`
                UserAns    string `json:"user_ans"`
            } `json:"user_answer"`
        }

        if err := c.BindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid request body",
            })
            return
        }

        var experienceEarned int
        for _, answer := range requestBody.UserAnswer {
            // Find the corresponding question
            for _, soal := range soals {
                if answer.ID == int(soal.Id) && answer.UserAns == soal.Correct_answer {
                    experienceEarned += 10
                    // Break out of the inner loop once a match is found
                    break
                }
            }
        }

        result := gin.H{
            "level":             level,
            "experience_earned": experienceEarned,
        }

        c.JSON(http.StatusOK, gin.H{
            "result": result,
        })
    } else {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid level parameter",
        })
    }
}

