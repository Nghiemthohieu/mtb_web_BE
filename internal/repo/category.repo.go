package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type CategoryRepo struct{}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{}
}

func (r *CategoryRepo) GetAll() ([]models.Category, int, error) {
	var categories []models.Category
	if err := global.Mdb.Preload("Parent").Preload("Products").Find(&categories).Error; err != nil {
		return nil, 20068, err
	}
	return categories, 20001, nil
}

func (r *CategoryRepo) GetById(id int) (models.Category, int, error) {
	var category models.Category
	if err := global.Mdb.Preload("Parent").Preload("Products").First(&category, id).Error; err != nil {
		return models.Category{}, 20067, err
	}
	return category, 20001, nil
}

func (r *CategoryRepo) Create(category *models.Category) (int, error) {
	if err := global.Mdb.Create(category).Error; err != nil {
		return 20062, err
	}
	return 20001, nil
}

func (r *CategoryRepo) Update(category *models.Category) (int, error) {
	var existing models.Category
	if err := global.Mdb.Where("id = ?", category.ID).First(&existing).Error; err != nil {
		return 20063, err
	}
	category.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(category).Error; err != nil {
		return 20063, err
	}
	return 20001, nil
}

func (r *CategoryRepo) DeleteById(id int) (int, error) {
	var category models.Category
	if err := global.Mdb.First(&category, id).Error; err != nil {
		return 20064, err
	}
	if err := global.Mdb.Delete(&category).Error; err != nil {
		return 20064, err
	}
	return 20001, nil
}

func (r *CategoryRepo) DeleteByIDs(ids []int) (int, error) {
	var categories []models.Category
	if err := global.Mdb.Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return 2025, err
	}
	if err := global.Mdb.Where("id IN ?", ids).Delete(&categories).Error; err != nil {
		return 20065, err
	}
	return 20001, nil
}

func (r *CategoryRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Category{}).Error; err != nil {
		return 20066, err
	}
	return 20001, nil
}
