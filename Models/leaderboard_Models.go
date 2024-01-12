package models

import "time"

type Leaderboard struct {
	Id         uint `gorm:"primaryKey"`
	UserID     int
	User       *User
	Experience int `gorm:"type:int;default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
