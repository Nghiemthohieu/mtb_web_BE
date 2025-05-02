package models

import "gorm.io/gorm"

type ProductReviewImg struct {
	gorm.Model

	ProductReviewID uint          `json:"product_review_id" gorm:"column:product_review_id;not null;index"`
	ProductReview   ProductReview `json:"product_review" gorm:"foreignKey:ProductReviewID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Image string `json:"image" gorm:"type:text;not null"`
}

func (p *ProductReviewImg) TableName() string {
	return "go_db_product_review_imgs"
}
