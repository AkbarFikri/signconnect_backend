package models

import "time"

type Leaderboard struct {
	Id         uint `gorm:"primaryKey"`
	UserId     int
	User       User `gorm:"foreignKey:UserId"`
	Experience int  `gorm:"type:int;default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
