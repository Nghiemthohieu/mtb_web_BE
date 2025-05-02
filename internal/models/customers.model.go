package models

import (
	"time"
)

type Customer struct {
	UserID    *int      `json:"user_id" gorm:"primaryKey;type:int;column:user_id"`
	User      *User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Address   string    `json:"address" gorm:"type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)"`
	Birthdate time.Time `json:"birthdate" gorm:"type:date"`
	Height    int       `json:"height" gorm:"type:int"`
	Weight    int       `json:"weight" gorm:"type:int"`
}

func (c *Customer) TableName() string {
	return "go_db_customers"
}
