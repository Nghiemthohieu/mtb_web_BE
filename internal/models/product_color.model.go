package models

import "gorm.io/gorm"

type ProductColor struct {
	gorm.Model
	Name  string `gorm:"column:name;type:varchar(100);not null" json:"color"`
	Color string `gorm:"column:color;type:varchar(250);not null" json:"color_img"`
}

func (pc *ProductColor) TableName() string {
	return "go_db_product_colors"
}
