package models

type Leveling struct {
    Id             int     `gorm:"primaryKey"`
    Level          int     `gorm:"type:int"`
    UserId         int
    User           User    `gorm:"foreignKey:UserId"`
    Status         string  `gorm:"type:varchar(15)"`
}