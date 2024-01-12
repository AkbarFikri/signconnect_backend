package models

type Leveling struct {
	Id     int `gorm:"primaryKey"`
	Level  int `gorm:"type:int"`
	UserID int
	Status string `gorm:"type:varchar(15)"`
}
