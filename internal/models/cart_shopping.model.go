package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   int        `json:"user_id" gorm:"column:user_id;not null"`
	User     User       `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CartItem []CartItem `json:"cart_item" gorm:"foreignKey:CartID"`
	Price    float64    `json:"price" gorm:"column:price;not null"`
}

func (Cart) TableName() string {
	return "go_db_carts"
}
