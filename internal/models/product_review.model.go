package models

import "gorm.io/gorm"

type ProductReview struct {
	gorm.Model

	ProductID uint    `json:"product_id" gorm:"column:product_id;not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	UserID uint `json:"user_id" gorm:"column:user_id;not null"`
	User   User `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Rating           int                `json:"rating" gorm:"not null"`
	Review           string             `json:"review" gorm:"type:text"`
	ProductReviewImg []ProductReviewImg `json:"product_review_imgs" gorm:"foreignKey:ProductReviewID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// 👇 Đây là phần thêm vào đúng chuẩn nè
	ParentID *uint          `json:"parent_id" gorm:"column:parent_id"` // khóa ngoại
	Parent   *ProductReview `json:"parent" gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (p *ProductReview) TableName() string {
	return "go_db_product_reviews"
}
