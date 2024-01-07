package models

import "gorm.io/gorm"

type Leaderboard struct {
    gorm.Model
    UserID      uint   `json:"user_id" gorm:"uniqueIndex"`
    Experience  int    `json:"experience"`
    User        User   `json:"user" gorm:"foreignKey:UserID"`
}

func NewLeaderboard(userID uint, experience int) *Leaderboard {
    return &Leaderboard{
        UserID:     userID,
        Experience: experience,
    }
}
