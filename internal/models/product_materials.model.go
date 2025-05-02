package models

import "gorm.io/gorm"

type ProductMaterial struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255);not null;unique" json:"name"`
}

func (pm *ProductMaterial) TableName() string {
	return "go_db_product_materials"
}
