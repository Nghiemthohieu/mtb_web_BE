package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name              string          `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description       string          `gorm:"column:description;type:text" json:"description"`
	Slug              string          `gorm:"column:slug;type:varchar(255);not null" json:"slug"`
	Size              []ProductSize   `gorm:"many2many:go_product_sizes;"`
	Color             []ProductColor  `gorm:"many2many:go_product_colors;"`
	Categories        []Category      `gorm:"many2many:go_db_product_categories;"`
	ProductStyleID    *uint           `gorm:"column:product_style_id;index" json:"product_style_id"` // Thay đổi thành *uint cho phép NULL
	ProductStyle      ProductStyle    `gorm:"foreignKey:ProductStyleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product_style"`
	ProductMaterialID uint            `gorm:"column:product_material_id" json:"product_material_id"`
	ProductMaterial   ProductMaterial `gorm:"foreignKey:ProductMaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product_material"`
	OriginalPrice     int             `gorm:"column:original_price;type:int;not null" json:"original_price"`
	DiscountPrice     int             `gorm:"column:discount_price;type:int;not null" json:"discount_price"`
	Price             int             `gorm:"column:price;type:int;not null" json:"price"`
	AvgRating         float32         `gorm:"column:avg_rating;type:float" json:"avg_rating"`
	RatingCount       int             `gorm:"column:rating_count;type:int" json:"rating_count"`
	ProductImages     []ProductImage  `gorm:"foreignKey:ProductID" json:"product_images"`
}

func (p *Product) TableName() string {
	return "go_db_products"
}
