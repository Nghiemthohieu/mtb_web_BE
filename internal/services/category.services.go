package services

import (
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	util "mtb_web/pkg/utils"
)

type CategoryService struct {
	CategoryRepo *repo.CategoryRepo
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		CategoryRepo: repo.NewCategoryRepo(),
	}
}

func (s *CategoryService) GetAll() ([]models.Category, int, error) {
	return s.CategoryRepo.GetAll()
}

func (s *CategoryService) GetByID(id int) (models.Category, int, error) {
	return s.CategoryRepo.GetById(id)
}

func (s *CategoryService) Create(category *models.Category) (int, error) {
	imgData, err := util.DecodeBase64Image(category.Img)
	if err != nil {
		return 20062, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20062, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	category.Img = imgURL
	return s.CategoryRepo.Create(category)
}

func (s *CategoryService) Update(category *models.Category) (int, error) {
	imgData, err := util.DecodeBase64Image(category.Img)
	if err != nil {
		return 20062, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20062, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	category.Img = imgURL
	return s.CategoryRepo.Update(category)
}

func (s *CategoryService) DeleteById(id int) (int, error) {
	return s.CategoryRepo.DeleteById(id)
}

func (s *CategoryService) DeleteByIDs(ids []int) (int, error) {
	return s.CategoryRepo.DeleteByIDs(ids)
}

func (s *CategoryService) DeleteAll() (int, error) {
	return s.CategoryRepo.DeleteAll()
}
