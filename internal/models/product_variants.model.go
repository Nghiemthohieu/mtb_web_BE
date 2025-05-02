package models

import "gorm.io/gorm"

type ProductVariant struct {
	gorm.Model

	ProductID uint    `json:"product_id" gorm:"column:product_id;not null;index:idx_product_size_color"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	SizeID uint        `json:"size_id" gorm:"column:size_id;not null;index:idx_product_size_color"`
	Size   ProductSize `json:"size" gorm:"foreignKey:SizeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ColorID uint         `json:"color_id" gorm:"column:color_id;not null;index:idx_product_size_color"`
	Color   ProductColor `json:"color" gorm:"foreignKey:ColorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Stock int `json:"stock" gorm:"not null"`
}

func (pv *ProductVariant) TableName() string {
	return "go_db_product_variants"
}
