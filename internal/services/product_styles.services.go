package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type ProductStylesService struct {
	ProductStylesRepo *repo.ProductStylesRepo
}

func NewProductStylesService() *ProductStylesService {
	return &ProductStylesService{
		ProductStylesRepo: repo.NewProductStylesRepo(),
	}
}

func (s *ProductStylesService) GetAll() ([]models.ProductStyle, int, error) {
	return s.ProductStylesRepo.GetAll()
}

func (s *ProductStylesService) GetByID(Id int) (models.ProductStyle, int, error) {
	return s.ProductStylesRepo.GetById(Id)
}

func (s *ProductStylesService) Create(ProductStyles *models.ProductStyle) (int, error) {
	return s.ProductStylesRepo.Create(ProductStyles)
}

func (s *ProductStylesService) Update(ProductStyle *models.ProductStyle) (int, error) {
	return s.ProductStylesRepo.Update(ProductStyle)
}

func (s *ProductStylesService) DeleteById(Id int) (int, error) {
	return s.ProductStylesRepo.DeleteById(Id)
}

func (s *ProductStylesService) DeleteByIDs(ids []int) (int, error) {
	return s.ProductStylesRepo.DeleteByIDs(ids)
}
func (s *ProductStylesService) DeleteAll() (int, error) {
	return s.ProductStylesRepo.DeleteAll()
}