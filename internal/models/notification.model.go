package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID  *uint      `json:"user_id" gorm:"column:user_id;not null;index"`
	User    *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Title   string     `json:"title" gorm:"column:title;not null"`
	Message string     `json:"message" gorm:"column:message;not null"`
	Type    string     `json:"type" gorm:"column:type;not null"`
	IsRead  bool       `json:"is_read" gorm:"column:is_read;default:false"`
	ReadAt  *time.Time `json:"read_at,omitempty" gorm:"column:read_at"`
}

func (n *Notification) TableName() string {
	return "go_db_notifications"
}
