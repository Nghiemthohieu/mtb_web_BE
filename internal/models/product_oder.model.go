package models

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model

	OrderID *uint `json:"order_id" gorm:"column:order_id"` // allow NULL
	Order   Order `json:"order" gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	ProductID *uint   `json:"product_id" gorm:"column:product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	SizeID *uint       `json:"size_id" gorm:"column:size_id"`
	Size   ProductSize `json:"size" gorm:"foreignKey:SizeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	ColorID *uint        `json:"color_id" gorm:"column:color_id"`
	Color   ProductColor `json:"color" gorm:"foreignKey:ColorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Quantity int `json:"quantity" gorm:"not null"`
}

func (po *ProductOrder) TableName() string {
	return "go_db_product_orders"
}
