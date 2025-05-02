package models

import "gorm.io/gorm"

type SlideShow struct {
	gorm.Model
	Image string `json:"image" gorm:"column:image;type:varchar(255);not null"`
	Title string `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Slug  string `gorm:"column:slug;type:varchar(255);not null" json:"slug"`
}

func (SlideShow) TableName() string {
	return "go_db_slide_shows"
}
