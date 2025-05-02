package models

import "gorm.io/gorm"

type ProductSize struct {
	gorm.Model
	Size string `gorm:"column:size;type:varchar(50);not null;unique" json:"size"`
}

func (ps *ProductSize) TableName() string {
	return "go_db_product_sizes"
}
