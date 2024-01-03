package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)" json:"username"`
	Email    string `gorm:"uniqe" json:"email"`
	Password string `json:"password"`
}
