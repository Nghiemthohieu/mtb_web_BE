package models

import "time"

type Staffs struct {
	UserID    uint      `gorm:"primaryKey;column:user_id;type:int unsigned" json:"user_id"` // Đổi thành 'uint' để phù hợp với kiểu 'INT UNSIGNED'
	User      *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`
	Name      string    `gorm:"column:name" json:"name"`
	Position  string    `gorm:"column:position" json:"position"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Status    string    `gorm:"column:status;type:varchar(50);not null;default:'active'" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (s *Staffs) TableName() string {
	return "go_db_staffs"
}
