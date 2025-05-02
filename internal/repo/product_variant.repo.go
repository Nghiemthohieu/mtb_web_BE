package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type ProductVariantRepo struct{}

func NewProductVariantRepo() *ProductVariantRepo {
	return &ProductVariantRepo{}
}

func (r *ProductVariantRepo) GetAll() ([]models.ProductVariant, int, error) {
	var variants []models.ProductVariant
	err := global.Mdb.Preload("Product").
		Preload("Size").
		Preload("Color").
		Find(&variants).Error
	if err != nil {
		return nil, 20078, err
	}
	return variants, 20001, nil
}

func (r *ProductVariantRepo) GetById(id int) (models.ProductVariant, int, error) {
	var variant models.ProductVariant
	err := global.Mdb.Preload("Product").
		Preload("Size").
		Preload("Color").
		First(&variant, id).Error
	if err != nil {
		return models.ProductVariant{}, 20077, err
	}
	return variant, 20001, nil
}

func (r *ProductVariantRepo) Create(variant *models.ProductVariant) (int, error) {
	err := global.Mdb.Create(variant).Error
	if err != nil {
		return 20072, err
	}
	return 20001, nil
}

func (r *ProductVariantRepo) Update(variant *models.ProductVariant) (int, error) {
	var existing models.ProductVariant
	if err := global.Mdb.First(&existing, variant.ID).Error; err != nil {
		return 20073, err
	}
	variant.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(variant).Error; err != nil {
		return 20073, err
	}
	return 20001, nil
}

func (r *ProductVariantRepo) DeleteById(id int) (int, error) {
	var variant models.ProductVariant
	if err := global.Mdb.First(&variant, id).Error; err != nil {
		return 20074, err
	}
	if err := global.Mdb.Delete(&variant).Error; err != nil {
		return 20074, err
	}
	return 20001, nil
}

func (r *ProductVariantRepo) DeleteByIDs(ids []int) (int, error) {
	var variants []models.ProductVariant
	if err := global.Mdb.Where("id IN ?", ids).Find(&variants).Error; err != nil {
		return 2075, err
	}
	if err := global.Mdb.Where("id IN ?", ids).Delete(&variants).Error; err != nil {
		return 20075, err
	}
	return 20001, nil
}

func (r *ProductVariantRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.ProductVariant{}).Error; err != nil {
		return 20076, err
	}
	return 20001, nil
}
