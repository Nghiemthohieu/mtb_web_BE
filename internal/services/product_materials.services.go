package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type ProductMaterialsService struct {
	ProductMaterialsRepo *repo.ProductMaterialsRepo
}

func NewProductMaterialsService() *ProductMaterialsService {
	return &ProductMaterialsService{
		ProductMaterialsRepo: repo.NewProductMaterialsRepo(),
	}
}

func (s *ProductMaterialsService) GetAll() ([]models.ProductMaterial, int, error) {
	return s.ProductMaterialsRepo.GetAll()
}

func (s *ProductMaterialsService) GetByID(id int) (models.ProductMaterial, int, error) {
	return s.ProductMaterialsRepo.GetById(id)
}

func (s *ProductMaterialsService) Create(material *models.ProductMaterial) (int, error) {
	return s.ProductMaterialsRepo.Create(material)
}

func (s *ProductMaterialsService) Update(material *models.ProductMaterial) (int, error) {
	return s.ProductMaterialsRepo.Update(material)
}

func (s *ProductMaterialsService) DeleteById(id int) (int, error) {
	return s.ProductMaterialsRepo.DeleteById(id)
}

func (s *ProductMaterialsService) DeleteByIDs(ids []int) (int, error) {
	return s.ProductMaterialsRepo.DeleteByIDs(ids)
}

func (s *ProductMaterialsService) DeleteAll() (int, error) {
	return s.ProductMaterialsRepo.DeleteAll()
}
