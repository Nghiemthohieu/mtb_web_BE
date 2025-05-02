package repo

import (
	"errors"
	"fmt"
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type ProductSizeRepo struct{}

func NewProductSizeRepo() *ProductSizeRepo {
	return &ProductSizeRepo{}
}

func (r *ProductSizeRepo) GetAll() ([]models.ProductSize, int, error) {
	var productSizes []models.ProductSize
	err := global.Mdb.Find(&productSizes).Error
	if err != nil {
		return nil, 20018, err
	}
	return productSizes, 20001, nil
}

func (r *ProductSizeRepo) GetByID(id uint) (*models.ProductSize, int, error) {
	var productSize models.ProductSize
	err := global.Mdb.First(&productSize, id).Error
	if err != nil {
		return nil, 20017, err
	}
	return &productSize, 20001, nil
}

func (r *ProductSizeRepo) Create(productSize *models.ProductSize) (int, error) {
	var existing models.ProductSize
	err := global.Mdb.Unscoped().Where("size = ?", productSize.Size).First(&existing).Error
	if err == nil {
		if existing.DeletedAt.Valid {
			// Bản ghi đã bị soft delete → khôi phục
			existing.DeletedAt = gorm.DeletedAt{}
			err := global.Mdb.Unscoped().Save(&existing).Error
			if err != nil {
				// handle update error nếu cần
				return 20012, err
			}
			return 20001, nil
		} else {
			// Bản ghi đã tồn tại và không bị xoá → không cần thêm
			return 20012, errors.New("product size already exists")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Bản ghi chưa tồn tại → thêm mới
		err = global.Mdb.Create(&productSize).Error
		if err != nil {
			// handle create error nếu cần
			return 20012, err
		}
		return 20001, nil
	} else {
		// handle error nếu có lỗi khác
		return 20012, err
	}
}
func (r *ProductSizeRepo) Update(productSize *models.ProductSize) (int, error) {
	fmt.Println("productSize", productSize)
	var existingProductSize models.ProductSize
	err := global.Mdb.First(&existingProductSize, productSize.ID).Error
	if err != nil {
		return 20013, err
	}
	productSize.CreatedAt = existingProductSize.CreatedAt
	err = global.Mdb.Save(productSize).Error
	if err != nil {
		return 20013, err
	}
	return 20001, nil
}
func (r *ProductSizeRepo) Delete(id uint) (int, error) {
	var productSize models.ProductSize
	err := global.Mdb.First(&productSize, id).Error
	if err != nil {
		return 20016, err
	}
	err = global.Mdb.Delete(&productSize).Error
	if err != nil {
		return 20014, err
	}
	return 20001, nil
}
func (r *ProductSizeRepo) DeleteByIDs(ids []uint) (int, error) {
	var productSizes []models.ProductSize
	err := global.Mdb.Where("id IN ?", ids).Find(&productSizes).Error
	if err != nil {
		return 2019, err
	}
	if err := global.Mdb.Where("id IN ?", ids).Delete(&productSizes).Error; err != nil {
		return 20015, err
	}
	return 20001, nil
}
func (r *ProductSizeRepo) DeleteAll() (int, error) {
	var productSizes []models.ProductSize
	err := global.Mdb.Find(&productSizes).Error
	if err != nil {
		return 20018, err
	}
	err = global.Mdb.Delete(&productSizes).Error
	if err != nil {
		return 20016, err
	}
	return 20001, nil
}
