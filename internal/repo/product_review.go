package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"
)

type ProductReviewRepo struct{}

func NewProductReviewRepo() *ProductReviewRepo {
	return &ProductReviewRepo{}
}

// Lấy tất cả đánh giá sản phẩm
func (r *ProductReviewRepo) GetAll() ([]models.ProductReview, int, error) {
	var reviews []models.ProductReview
	if err := global.Mdb.
		Preload("Product").
		Preload("User").
		Preload("ProductReviewImg").
		Find(&reviews).Error; err != nil {
		return nil, 20028, err
	}
	return reviews, 20001, nil
}

// Lấy đánh giá theo ID sản phẩm
func (r *ProductReviewRepo) GetByProductID(productID uint) ([]models.ProductReview, int, error) {
	var reviews []models.ProductReview
	if err := global.Mdb.
		Preload("User").
		Preload("ProductReviewImg").
		Where("product_id = ?", productID).
		Find(&reviews).Error; err != nil {
		return nil, 20016, err
	}
	return reviews, 20001, nil
}

// Tạo đánh giá mới
func (r *ProductReviewRepo) Create(review *models.ProductReview) (int, error) {
	if err := global.Mdb.Create(review).Error; err != nil {
		return 20022, err
	}
	return 20001, nil
}

// Cập nhật đánh giá
func (r *ProductReviewRepo) Update(review *models.ProductReview) (int, error) {
	if err := global.Mdb.Save(review).Error; err != nil {
		return 20023, err
	}
	return 20001, nil
}

// Xóa đánh giá theo ID
func (r *ProductReviewRepo) DeleteByID(id uint) (int, error) {
	if err := global.Mdb.Delete(&models.ProductReview{}, id).Error; err != nil {
		return 20024, err
	}
	return 20001, nil
}
