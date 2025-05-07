package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    *uint    `json:"cart_id" gorm:"column:cart_id;not null"`
	Cart      *Cart    `json:"cart" gorm:"foreignKey:CartID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductID *uint    `json:"product_id" gorm:"column:product_id;not null"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Quantity  int      `json:"quantity" gorm:"type:int;not null"`
	Price     float64  `json:"price" gorm:"column:price;not null"`
}

func (CartItem) TableName() string {
	return "go_db_cart_Items"
}
