package services

import (
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	util "mtb_web/pkg/utils"
)

type ProductColorService struct {
	ProductColorRepo *repo.ProductColorRepo
}

func NewProductColorService() *ProductColorService {
	return &ProductColorService{
		ProductColorRepo: repo.NewProductColorRepo(),
	}
}
func (s *ProductColorService) Create(productColor *models.ProductColor) (int, error) {
	imgData, err := util.DecodeBase64Image(productColor.Color)
	if err != nil {
		return 20062, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20062, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	productColor.Color = imgURL
	return s.ProductColorRepo.Create(productColor)
}
func (s *ProductColorService) Update(productColor *models.ProductColor) (int, error) {
	imgData, err := util.DecodeBase64Image(productColor.Color)
	if err != nil {
		return 20062, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20062, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	productColor.Color = imgURL
	return s.ProductColorRepo.Update(productColor)
}
func (s *ProductColorService) GetById(id int) (*models.ProductColor, int, error) {
	return s.ProductColorRepo.GetById(id)
}
func (s *ProductColorService) GetAll() ([]*models.ProductColor, int, error) {
	return s.ProductColorRepo.GetAll()
}
func (s *ProductColorService) DeleteById(id int) (int, error) {
	return s.ProductColorRepo.DeleteById(id)
}
func (s *ProductColorService) DeleteByIDs(ids []int) (int, error) {
	return s.ProductColorRepo.DeleteByIDs(ids)
}
func (s *ProductColorService) DeleteAll() (int, error) {
	return s.ProductColorRepo.DeleteAll()
}
