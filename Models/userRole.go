package models

import "time"

type UserRole struct {
	Id         uint 		`gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserId     int
	User       User     	`gorm:"foreignKey:UserId"`
	Role       string   	`gorm:"default:"user";type:varchar(20)" json:role`
	Lembaga_id int      	`gorm:"default:0"`
	Lembaga    *Lembaga 	`gorm:"foreignKey:Lembaga_id"`
}
