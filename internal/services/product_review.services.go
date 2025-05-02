package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type ProductReviewService struct {
	ProductReviewRepo *repo.ProductReviewRepo
}

func NewProductReviewService() *ProductReviewService {
	return &ProductReviewService{
		ProductReviewRepo: repo.NewProductReviewRepo(),
	}
}

func (s *ProductReviewService) GetAll() ([]models.ProductReview, int, error) {
	return s.ProductReviewRepo.GetAll()
}

func (s *ProductReviewService) GetByProductID(productID uint) ([]models.ProductReview, int, error) {
	return s.ProductReviewRepo.GetByProductID(productID)
}

func (s *ProductReviewService) Create(review *models.ProductReview) (int, error) {
	return s.ProductReviewRepo.Create(review)
}

func (s *ProductReviewService) Update(review *models.ProductReview) (int, error) {
	return s.ProductReviewRepo.Update(review)
}

func (s *ProductReviewService) DeleteById(id uint) (int, error) {
	return s.ProductReviewRepo.DeleteByID(id)
}
