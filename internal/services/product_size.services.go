package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type ProductSizeService struct {
	ProductSizeRepo *repo.ProductSizeRepo
}

func NewProductSizeService() *ProductSizeService {
	return &ProductSizeService{
		ProductSizeRepo: repo.NewProductSizeRepo(),
	}
}
func (s *ProductSizeService) GetAll() ([]models.ProductSize, int, error) {
	return s.ProductSizeRepo.GetAll()
}
func (s *ProductSizeService) GetByID(id uint) (*models.ProductSize, int, error) {
	return s.ProductSizeRepo.GetByID(id)
}
func (s *ProductSizeService) Create(productSize *models.ProductSize) (int, error) {
	return s.ProductSizeRepo.Create(productSize)
}
func (s *ProductSizeService) Update(productSize *models.ProductSize) (int, error) {
	return s.ProductSizeRepo.Update(productSize)
}
func (s *ProductSizeService) Delete(id uint) (int, error) {
	return s.ProductSizeRepo.Delete(id)
}
func (s *ProductSizeService) DeleteByIDs(ids []uint) (int, error) {
	return s.ProductSizeRepo.DeleteByIDs(ids)
}

func (s *ProductSizeService) DeleteAll() (int, error) {
	return s.ProductSizeRepo.DeleteAll()
}
