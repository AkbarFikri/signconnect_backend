package models

type Lembaga struct {
    Id             int     `gorm:"primaryKey"`
    Nama           string  `gorm:"type:varchar(255)"`
    Tempat         string  `gorm:"type:varchar(255)"`
    Deskripsi      string  `gorm:"type:text"`
    Pekerjaan      string  `gorm:"type:text"`
    Persyaratan    string  `gorm:"type:text"`
    image_url      string  `gorm:"type:varchar(511)"`
    min_pengalaman int     `gorm:"type:int"`
    Translator     []UserRole `gorm:"onetomany"`
}