package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type PermissionRepo struct{}

func NewPermissionRepo() *PermissionRepo {
	return &PermissionRepo{}
}

func (r *PermissionRepo) GetAll() ([]models.Permission, int, error) {
	var permissions []models.Permission
	if err := global.Mdb.Find(&permissions).Error; err != nil {
		return nil, 20128, err
	}
	return permissions, 20001, nil
}

func (r *PermissionRepo) GetById(id int) (models.Permission, int, error) {
	var permission models.Permission
	if err := global.Mdb.First(&permission, id).Error; err != nil {
		return models.Permission{}, 20016, err
	}
	return permission, 20127, nil
}

func (r *PermissionRepo) Create(permission *models.Permission) (int, error) {
	if err := global.Mdb.Create(permission).Error; err != nil {
		return 20122, err
	}
	return 20001, nil
}

func (r *PermissionRepo) Update(permission *models.Permission) (int, error) {
	var existing models.Permission
	if err := global.Mdb.First(&existing, permission.ID).Error; err != nil {
		return 20123, err
	}
	permission.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(permission).Error; err != nil {
		return 20123, err
	}
	return 20001, nil
}

func (r *PermissionRepo) DeleteById(id int) (int, error) {
	var permission models.Permission
	if err := global.Mdb.First(&permission, id).Error; err != nil {
		return 20124, err
	}
	if err := global.Mdb.Delete(&permission).Error; err != nil {
		return 20124, err
	}
	return 20001, nil
}

func (r *PermissionRepo) DeleteByIDs(ids []int) (int, error) {
	var permissions []models.Permission
	if err := global.Mdb.Where("id IN ?", ids).Find(&permissions).Error; err != nil {
		return 20125, err
	}
	if err := global.Mdb.Delete(&permissions).Error; err != nil {
		return 20125, err
	}
	return 20001, nil
}

func (r *PermissionRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Permission{}).Error; err != nil {
		return 20126, err
	}
	return 20001, nil
}
