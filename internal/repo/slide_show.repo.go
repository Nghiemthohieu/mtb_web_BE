package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type SlideShowRepo struct{}

func NewSlideShowRepo() *SlideShowRepo {
	return &SlideShowRepo{}
}

// Lấy tất cả SlideShow
func (r *SlideShowRepo) GetAll() ([]models.SlideShow, int, error) {
	var slideshows []models.SlideShow
	if err := global.Mdb.Find(&slideshows).Error; err != nil {
		return nil, 20178, err
	}
	return slideshows, 20001, nil
}

// Lấy SlideShow theo ID
func (r *SlideShowRepo) GetByID(id uint) (*models.SlideShow, int, error) {
	var slideshow models.SlideShow
	if err := global.Mdb.First(&slideshow, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 20177, nil
		}
		return nil, 20177, err
	}
	return &slideshow, 20001, nil
}

// Tạo mới SlideShow
func (r *SlideShowRepo) Create(slideshow *models.SlideShow) (int, error) {
	if err := global.Mdb.Create(slideshow).Error; err != nil {
		return 20172, err
	}
	return 20001, nil
}

// Cập nhật SlideShow
func (r *SlideShowRepo) Update(slideshow *models.SlideShow) (int, error) {
	if err := global.Mdb.Save(slideshow).Error; err != nil {
		return 20173, err
	}
	return 20001, nil
}

// Xóa SlideShow theo ID
func (r *SlideShowRepo) DeleteByID(id uint) (int, error) {
	if err := global.Mdb.Delete(&models.SlideShow{}, id).Error; err != nil {
		return 20174, err
	}
	return 20001, nil
}
