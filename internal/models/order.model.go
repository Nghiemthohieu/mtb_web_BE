package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          int            `gorm:"column:user_id" json:"user_id"`
	User            User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`
	Name            string         `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Email           string         `gorm:"column:email;type:varchar(100);not null" json:"email"`
	Phone           string         `gorm:"column:phone;type:varchar(100);not null" json:"phone"`
	Address         string         `gorm:"column:address;type:varchar(100);not null" json:"address"`
	TotalPrice      float64        `gorm:"column:total_price;not null" json:"total_price"`
	PaymentMethodID *uint          `gorm:"column:payment_method_id;index" json:"payment_method_id"`
	PaymentMethod   PaymentMethod  `gorm:"foreignKey:PaymentMethodID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"payment_method"`
	ProductOrders   []ProductOrder `gorm:"foreignKey:OrderID" json:"product_orders"`
	StatusID        *uint          `gorm:"column:status_id" json:"status_id"` // Thay đổi thành *uint để cho phép NULL
	Status          OrderStatus    `gorm:"foreignKey:StatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"status"`
}

func (o *Order) TableName() string {
	return "go_db_orders"
}
