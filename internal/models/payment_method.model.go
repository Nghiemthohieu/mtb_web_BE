package models

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name;type:varchar(255);not null;unique"`
	Code        string `json:"code" gorm:"column:code;type:varchar(255);not null;unique"`
	Description string `json:"description" gorm:"column:description;type:text"`
	Enabled     bool   `json:"enabled" gorm:"column:enabled;type:bool;default:true"`

	// Optional: Liên kết với Orders nếu cần
	// Orders      []Order `gorm:"foreignKey:PaymentMethodID" json:"orders,omitempty"`
}

func (p *PaymentMethod) TableName() string {
	return "go_db_payment_methods"
}
