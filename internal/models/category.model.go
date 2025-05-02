package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Slug     string    `gorm:"column:slug;type:varchar(255);not null" json:"slug"`
	Img      string    `gorm:"column:img;type:varchar(255);not null" json:"img"`
	ParentId *uint     `gorm:"column:parent_id" json:"parent_id"` // Cho phép NULL
	Parent   *Category `gorm:"foreignKey:ParentId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"parent"`
	Products []Product `gorm:"many2many:go_db_product_categories;"`
}

func (c *Category) TableName() string {
	return "go_db_categories"
}
