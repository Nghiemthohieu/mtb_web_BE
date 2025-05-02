package repo

import (
	"errors"
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type ProductStylesRepo struct{}

func NewProductStylesRepo() *ProductStylesRepo {
	return &ProductStylesRepo{}
}

func (r *ProductStylesRepo) GetAll() ([]models.ProductStyle, int, error) {
	var productStyles []models.ProductStyle
	if err := global.Mdb.Model(&productStyles).Find(&productStyles).Error; err != nil {
		return nil, 20028, err
	}
	return productStyles, 20001, nil
}

func (r *ProductStylesRepo) GetById(id int) (models.ProductStyle, int, error) {
	var productStyle models.ProductStyle
	if err := global.Mdb.Model(&productStyle).Where("id = ?", id).First(&productStyle).Error; err != nil {
		return models.ProductStyle{}, 20027, err
	}
	return productStyle, 20001, nil
}
func (r *ProductStylesRepo) Create(productStyle *models.ProductStyle) (int, error) {
	var existing models.ProductStyle
	err := global.Mdb.Unscoped().Where("name = ?", productStyle.Name).First(&existing).Error
	if err == nil {
		if existing.DeletedAt.Valid {
			// Bản ghi đã bị soft delete → khôi phục
			existing.DeletedAt = gorm.DeletedAt{}
			err := global.Mdb.Unscoped().Save(&existing).Error
			if err != nil {
				// handle update error nếu cần
				return 20022, err
			}
			return 20001, nil
		} else {
			// Bản ghi đã tồn tại và không bị xoá → không cần thêm
			return 20022, errors.New("product size already exists")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Bản ghi chưa tồn tại → thêm mới
		err = global.Mdb.Create(&productStyle).Error
		if err != nil {
			// handle create error nếu cần
			return 20022, err
		}
		return 20001, nil
	} else {
		// handle error nếu có lỗi khác
		return 20022, err
	}
}
func (r *ProductStylesRepo) DeleteById(id int) (int, error) {
	var productStyle models.ProductStyle
	err := global.Mdb.First(&productStyle, id).Error
	if err != nil {
		return 20016, err
	}
	if err := global.Mdb.Where("id = ?", id).Delete(&productStyle).Error; err != nil {
		return 20024, err
	}
	return 20001, nil
}
func (r *ProductStylesRepo) DeleteByIDs(ids []int) (int, error) {
	var productStyle []models.ProductStyle
	err := global.Mdb.Where("id IN ?", ids).Find(&productStyle).Error
	if err != nil {
		return 2019, err
	}
	if err := global.Mdb.Where("id IN ?", ids).Delete(&productStyle).Error; err != nil {
		return 20025, err
	}
	return 20001, nil
}
func (r *ProductStylesRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Delete(&models.ProductStyle{}).Error; err != nil {
		return 20026, err
	}
	return 20001, nil
}

func (r *ProductStylesRepo) Update(productStyle *models.ProductStyle) (int, error) {
	var existingProductStyle models.ProductStyle
	if err := global.Mdb.Where("id = ?", productStyle.ID).First(&existingProductStyle).Error; err != nil {
		return 20023, err
	}
	productStyle.CreatedAt = existingProductStyle.CreatedAt
	if err := global.Mdb.Save(productStyle).Error; err != nil {
		return 20023, err
	}
	return 20001, nil
}
