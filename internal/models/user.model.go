package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string    `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
	Password   *string   `json:"password,omitempty" gorm:"column:password;type:varchar(255)"`       // nullable
	Provider   *string   `json:"provider,omitempty" gorm:"column:provider;type:varchar(50)"`        // "google", "facebook", "local"
	ProviderID *string   `json:"provider_id,omitempty" gorm:"column:provider_id;type:varchar(255)"` // ID từ Google/Facebook
	Role       []Role    `json:"role" gorm:"many2many:go_db_user_roles;"`
	Customer   *Customer `json:"customer" gorm:"foreignKey:UserID"`
	Staff      *Staffs   `json:"staff" gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "go_db_users"
}
