package models

import "time"

type UserRole struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      *User
	Role      string   `gorm:"default:'user';type:varchar(20)" json:"role"`
	LembagaID int      `gorm:"default:0"`
	Lembaga   *Lembaga `gorm:"foreignKey:LembagaID"`
}
