package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type PermissionService struct {
	PermissionRepo *repo.PermissionRepo
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		PermissionRepo: repo.NewPermissionRepo(),
	}
}

func (s *PermissionService) GetAll() ([]models.Permission, int, error) {
	return s.PermissionRepo.GetAll()
}

func (s *PermissionService) GetByID(id int) (models.Permission, int, error) {
	return s.PermissionRepo.GetById(id)
}

func (s *PermissionService) Create(permission *models.Permission) (int, error) {
	return s.PermissionRepo.Create(permission)
}

func (s *PermissionService) Update(permission *models.Permission) (int, error) {
	return s.PermissionRepo.Update(permission)
}

func (s *PermissionService) DeleteById(id int) (int, error) {
	return s.PermissionRepo.DeleteById(id)
}

func (s *PermissionService) DeleteByIDs(ids []int) (int, error) {
	return s.PermissionRepo.DeleteByIDs(ids)
}

func (s *PermissionService) DeleteAll() (int, error) {
	return s.PermissionRepo.DeleteAll()
}
