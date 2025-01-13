package main

type URLMapping struct {
	ID       uint   `gorm:"primaryKey"`
	ShortURL string `gorm:"uniqueIndex"`
	LongURL  string `gorm:"not null"`
}
