package repo

import (
	"errors"
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type ProductMaterialsRepo struct{}

func NewProductMaterialsRepo() *ProductMaterialsRepo {
	return &ProductMaterialsRepo{}
}

func (r *ProductMaterialsRepo) GetAll() ([]models.ProductMaterial, int, error) {
	var materials []models.ProductMaterial
	if err := global.Mdb.Model(&materials).Find(&materials).Error; err != nil {
		return nil, 20048, err
	}
	return materials, 20001, nil
}

func (r *ProductMaterialsRepo) GetById(id int) (models.ProductMaterial, int, error) {
	var material models.ProductMaterial
	if err := global.Mdb.Model(&material).Where("id = ?", id).First(&material).Error; err != nil {
		return models.ProductMaterial{}, 200047, err
	}
	return material, 20001, nil
}

func (r *ProductMaterialsRepo) Create(material *models.ProductMaterial) (int, error) {
	var existing models.ProductMaterial
	err := global.Mdb.Unscoped().Where("name = ?", material.Name).First(&existing).Error
	if err == nil {
		if existing.DeletedAt.Valid {
			// Bản ghi bị soft delete → khôi phục
			existing.DeletedAt = gorm.DeletedAt{}
			if err := global.Mdb.Unscoped().Save(&existing).Error; err != nil {
				return 20042, err
			}
			return 20001, nil
		}
		return 20022, errors.New("product material already exists")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := global.Mdb.Create(material).Error; err != nil {
			return 20042, err
		}
		return 20001, nil
	} else {
		return 20042, err
	}
}

func (r *ProductMaterialsRepo) Update(material *models.ProductMaterial) (int, error) {
	var existing models.ProductMaterial
	if err := global.Mdb.Where("id = ?", material.ID).First(&existing).Error; err != nil {
		return 20043, err
	}
	material.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(material).Error; err != nil {
		return 20043, err
	}
	return 20001, nil
}

func (r *ProductMaterialsRepo) DeleteById(id int) (int, error) {
	var material models.ProductMaterial
	if err := global.Mdb.First(&material, id).Error; err != nil {
		return 20046, err
	}
	if err := global.Mdb.Where("id = ?", id).Delete(&material).Error; err != nil {
		return 20044, err
	}
	return 20001, nil
}

func (r *ProductMaterialsRepo) DeleteByIDs(ids []int) (int, error) {
	var materials []models.ProductMaterial
	if err := global.Mdb.Where("id IN ?", ids).Find(&materials).Error; err != nil {
		return 2049, err
	}
	if err := global.Mdb.Where("id IN ?", ids).Delete(&materials).Error; err != nil {
		return 20045, err
	}
	return 20001, nil
}

func (r *ProductMaterialsRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Delete(&models.ProductMaterial{}).Error; err != nil {
		return 20046, err
	}
	return 20001, nil
}
