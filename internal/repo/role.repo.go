package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type RoleRepo struct{}

func NewRoleRepo() *RoleRepo {
	return &RoleRepo{}
}

func (r *RoleRepo) GetAll() ([]models.Role, int, error) {
	var roles []models.Role
	if err := global.Mdb.Preload("Permissions").Find(&roles).Error; err != nil {
		return nil, 20118, err
	}
	return roles, 20001, nil
}

func (r *RoleRepo) GetById(id int) (models.Role, int, error) {
	var role models.Role
	if err := global.Mdb.Preload("Permissions").First(&role, id).Error; err != nil {
		return models.Role{}, 20016, err
	}
	return role, 20117, nil
}

func (r *RoleRepo) Create(role *models.Role) (int, error) {
	if err := global.Mdb.Create(role).Error; err != nil {
		return 20112, err
	}
	return 20001, nil
}

func (r *RoleRepo) Update(role *models.Role) (int, error) {
	var existing models.Role
	if err := global.Mdb.First(&existing, role.ID).Error; err != nil {
		return 20113, err
	}
	role.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Session(&gorm.Session{FullSaveAssociations: true}).Save(role).Error; err != nil {
		return 20113, err
	}
	return 20001, nil
}

func (r *RoleRepo) DeleteById(id int) (int, error) {
	var role models.Role
	if err := global.Mdb.First(&role, id).Error; err != nil {
		return 20115, err
	}
	if err := global.Mdb.Select("Permissions").Delete(&role).Error; err != nil {
		return 20114, err
	}
	return 20001, nil
}

func (r *RoleRepo) DeleteByIDs(ids []int) (int, error) {
	var roles []models.Role
	if err := global.Mdb.Where("id IN ?", ids).Find(&roles).Error; err != nil {
		return 20115, err
	}
	if err := global.Mdb.Select("Permissions").Delete(&roles).Error; err != nil {
		return 20115, err
	}
	return 20001, nil
}

func (r *RoleRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Role{}).Error; err != nil {
		return 20116, err
	}
	return 20001, nil
}
