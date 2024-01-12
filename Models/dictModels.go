package models

type DictionaryKategori struct {
	Id       int    `gorm:"primaryKey"`
	Kategori string `json:"kategori"`
}

type Dictionary struct {
	Id         int    `gorm:"primaryKey"`
	KategoriID uint   `json:"kategori_id"`
	Huruf      string `json:"huruf"`
	ImageURL   string `json:"image_url"`
}