# SIGNCONNECT DATABASE SPECIFICATION

## Table User

```golang
type User struct {
    gorm.model()
    Username    string  `gorm:"type:varchar(255)"` `json:username`
    Email       string  `gorm:"unique;type:varchar(255)"` `json:email`
    Password    string  `gorm:"type:text"` `json:password`
    Role        *UserRole
    Leveling    []models.Leveling `gorm:"onetomany"`
    Leaderboard *Leaderboard
}
```

## Table Soal

```golang
type Soal struct {
    Id              int     `gorm:"primaryKey;type:int"`
    QuestionImg_url string  `gorm:"type:varchar(511)"`
    Question        string  `gorm:"type:text"`
    answer_1        string  `gorm:"type:varchar(511)"`
    answer_2        string  `gorm:"type:varchar(511)"`
    answer_3        string  `gorm:"type:varchar(511)"`
    answer_4        string  `gorm:"type:varchar(511)"`
    correct_answer  string  `gorm:"type:varchar(511)"`
    Level           int     `gorm:"type:int"`
}
```

## Table Leveling

```golang
type Leveling struct {
    Id             int     `gorm:"primaryKey"`
    Level          int     `gorm:"type:int"`
    UserId         int
    User           User    `gorm:"foreignKey:UserId"`
    Status         string  `gorm:"type:varchar(15)"`
}
```

## Table Lembaga

```golang
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
```

## Table Leaderboard

```golang
type Leaderboard struct {
    Id          uint    `gorm:"primaryKey"`
    UserId      int
    User        User    `gorm:"foreignKey:UserId"`
    Experience  int     `gorm:"type:int;default:0"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## Tabel Role

```golang
type UserRole struct {
    Id          uint    `gorm:"primaryKey"`
    UserId      int
    User        User    `gorm:"foreignKey:UserId"`
    Experience  int     `gorm:"type:int;default:0"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```
