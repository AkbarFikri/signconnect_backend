package models

type Lembaga struct {
	Id             int        `gorm:"primaryKey"`
	Nama           string     `gorm:"type:varchar(255)"`
	Tempat         string     `gorm:"type:varchar(255)"`
	Deskripsi      string     `gorm:"type:text"`
	Pekerjaan      string     `gorm:"type:text"`
	Persyaratan    string     `gorm:"type:text"`
	Image_url      string     `gorm:"type:varchar(511)"`
	Min_pengalaman int        `gorm:"type:int"`
	Translator     []UserRole `gorm:"onetomany"`
}

type Lamaran struct {
	Id        int    `gorm:"primaryKey"`
	CvURL     string `gorm:"type:varchar(511)"`
	Username  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255)"`
	LembagaID int    `gorm:"type:int"`
	UserID    int    `gorm:"type:int"` // Foreign key to User model
	User      User   `gorm:"foreignKey:UserID"`
}
