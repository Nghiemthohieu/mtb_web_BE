package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"column:name;type:varchar(50);not null;unique" json:"name"`
	Description string       `gorm:"column:description;type:text" json:"description"` // Thêm mô tả nếu cần
	Permissions []Permission `gorm:"many2many:go_db_role_permissions;"`               // Chỉ rõ bảng liên kết
}

func (r *Role) TableName() string {
	return "go_db_roles"
}
