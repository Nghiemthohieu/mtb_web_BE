package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type ProductVariantService struct {
	ProductVariantRepo *repo.ProductVariantRepo
}

func NewProductVariantService() *ProductVariantService {
	return &ProductVariantService{
		ProductVariantRepo: repo.NewProductVariantRepo(),
	}
}

func (s *ProductVariantService) GetAll() ([]models.ProductVariant, int, error) {
	return s.ProductVariantRepo.GetAll()
}

func (s *ProductVariantService) GetByID(id int) (models.ProductVariant, int, error) {
	return s.ProductVariantRepo.GetById(id)
}

func (s *ProductVariantService) Create(variant *models.ProductVariant) (int, error) {
	return s.ProductVariantRepo.Create(variant)
}

func (s *ProductVariantService) Update(variant *models.ProductVariant) (int, error) {
	return s.ProductVariantRepo.Update(variant)
}

func (s *ProductVariantService) DeleteById(id int) (int, error) {
	return s.ProductVariantRepo.DeleteById(id)
}

func (s *ProductVariantService) DeleteByIDs(ids []int) (int, error) {
	return s.ProductVariantRepo.DeleteByIDs(ids)
}

func (s *ProductVariantService) DeleteAll() (int, error) {
	return s.ProductVariantRepo.DeleteAll()
}
