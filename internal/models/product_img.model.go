package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductID *uint   `gorm:"column:product_id" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product"`

	ImageURL string `gorm:"column:image_url;type:varchar(255);not null" json:"image_url"`
	IsMain   bool   `gorm:"column:is_main;default:false" json:"is_main"`

	ColorID     *uint        `gorm:"column:color_id" json:"color_id"`
	Color       ProductColor `gorm:"foreignKey:ColorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"color"`
	ImageBase64 string       `gorm:"-" json:"image_base64"` // Field này không lưu vào DB
}

func (pi *ProductImage) TableName() string {
	return "go_db_product_images"
}
