package services

import (
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	util "mtb_web/pkg/utils"
)

type SlideShowService struct {
	SlideShowRepo *repo.SlideShowRepo
}

func NewSlideShowService() *SlideShowService {
	return &SlideShowService{
		SlideShowRepo: repo.NewSlideShowRepo(),
	}
}

func (s *SlideShowService) GetAll() ([]models.SlideShow, int, error) {
	return s.SlideShowRepo.GetAll()
}

func (s *SlideShowService) GetByID(id int) (*models.SlideShow, int, error) {
	return s.SlideShowRepo.GetByID(uint(id))
}

func (s *SlideShowService) Create(slideshow *models.SlideShow) (int, error) {
	imgData, err := util.DecodeBase64Image(slideshow.Image)
	if err != nil {
		return 20062, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20062, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	slideshow.Image = imgURL
	return s.SlideShowRepo.Create(slideshow)
}

func (s *SlideShowService) Update(slideshow *models.SlideShow) (int, error) {
	imgData, err := util.DecodeBase64Image(slideshow.Image)
	if err != nil {
		return 20062, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20062, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	slideshow.Image = imgURL
	return s.SlideShowRepo.Update(slideshow)
}

func (s *SlideShowService) DeleteById(id int) (int, error) {
	return s.SlideShowRepo.DeleteByID(uint(id))
}
