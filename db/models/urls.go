package models

import "gorm.io/gorm"

const UrlsModel string = "Urls"

type Urls struct {
	gorm.Model
	ShortCode string `gorm:"uniqueIndex"`
	LongUrl   string `gorm:"uniqueIndex"`
}
