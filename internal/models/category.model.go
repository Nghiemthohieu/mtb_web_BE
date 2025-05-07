package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Slug     string    `gorm:"column:slug;type:varchar(255);not null" json:"slug"`
	Img      string    `gorm:"column:img;type:varchar(255);not null" json:"img"`
	ParentId *uint     `gorm:"column:parent_id" json:"parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"parent"`

	NameNavbar   string     `gorm:"column:name_navbar;type:varchar(255)" json:"name_navbar"`
	NameHome     string     `gorm:"column:name_home;type:varchar(255)" json:"name_home"`
	TitleDefault string     `gorm:"column:title_default;type:varchar(255)" json:"title_default"`
	TitleNavbar  string     `gorm:"column:title_navbar;type:varchar(255)" json:"title_navbar"`
	TitleHome    string     `gorm:"column:title_home;type:varchar(255)" json:"title_home"`
	IsVirtual    bool       `gorm:"column:is_virtual;default:false" json:"is_virtual"`
	Children     []Category `gorm:"many2many:go_db_virtual_categories;joinForeignKey:VirtualID;joinReferences:ChildID" json:"children"`
	IsNavbar     bool       `gorm:"column:is_navbar;default:false" json:"is_navbar"`
	IsFeatured   bool       `gorm:"column:is_featured;default:false" json:"is_featured"` // Hiển thị ở trang chủ?
	HomePosition int        `gorm:"column:home_position;default:0" json:"home_position"`

	Products []Product `gorm:"many2many:go_db_product_categories;" json:"products"`
}

func (c *Category) TableName() string {
	return "go_db_categories"
}
