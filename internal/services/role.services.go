package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type RoleService struct {
	RoleRepo *repo.RoleRepo
}

func NewRoleService() *RoleService {
	return &RoleService{
		RoleRepo: repo.NewRoleRepo(),
	}
}

func (s *RoleService) GetAll() ([]models.Role, int, error) {
	return s.RoleRepo.GetAll()
}

func (s *RoleService) GetByID(id int) (models.Role, int, error) {
	return s.RoleRepo.GetById(id)
}

func (s *RoleService) Create(role *models.Role) (int, error) {
	return s.RoleRepo.Create(role)
}

func (s *RoleService) Update(role *models.Role) (int, error) {
	return s.RoleRepo.Update(role)
}

func (s *RoleService) DeleteById(id int) (int, error) {
	return s.RoleRepo.DeleteById(id)
}

func (s *RoleService) DeleteByIDs(ids []int) (int, error) {
	return s.RoleRepo.DeleteByIDs(ids)
}

func (s *RoleService) DeleteAll() (int, error) {
	return s.RoleRepo.DeleteAll()
}
