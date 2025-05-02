package models

import "gorm.io/gorm"

type OrderStatus struct {
	gorm.Model
	Status string `json:"status" gorm:"column:status;type:varchar(100);unique;not null"`
}

func (o *OrderStatus) TableName() string {
	return "go_db_order_status"
}
