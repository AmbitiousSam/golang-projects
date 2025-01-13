package main

type URLMapping struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	shortURL string `gorm:"uniqueIndex" json:"short_url"`
	LongURL  string `json: long_url`
}
