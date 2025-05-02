package models

import "gorm.io/gorm"

type ProductStyle struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);not null;unique" json:"name"`
}

func (ps *ProductStyle) TableName() string {
	return "go_db_product_styles"
}
