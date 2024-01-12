package models

type Soal struct {
    Id              uint    `gorm:"primaryKey;type:uint"`
    QuestionImg_url string  `gorm:"type:varchar(511)"`
    Question        string  `gorm:"type:text"`
    Answer_1        string  `gorm:"type:varchar(511)"`
    Answer_2        string  `gorm:"type:varchar(511)"`
    Answer_3        string  `gorm:"type:varchar(511)"`
    Answer_4        string  `gorm:"type:varchar(511)"`
    Correct_answer  string  `gorm:"type:varchar(511)"`
    Level           int     `gorm:"type:int"`
}
