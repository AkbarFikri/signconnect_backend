package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(255)" json:"username"`
	Email       string `gorm:"unique;type:varchar(255)" json:"email"`
	Password    string `gorm:"type:text" json:"password"`
	Role        *UserRole
	Leveling    []Leveling `gorm:"onetomany"`
	Leaderboard *Leaderboard
}
