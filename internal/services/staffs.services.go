package services

import (
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	util "mtb_web/pkg/utils"
)

type StaffService struct {
	repo repo.StaffRepo
}

func NewStaffService() *StaffService {
	return &StaffService{
		repo: *repo.NewStaffRepo(),
	}
}

func (s *StaffService) GetAllStaff() ([]models.User, int, error) {
	return s.repo.GetAllStaff()
}

func (s *StaffService) GetStaffById(id int) (models.User, int, error) {
	return s.repo.GetStaffById(id)
}

func (s *StaffService) GetStaffByEmail(email string) (models.User, int, error) {
	return s.repo.GetStaffByEmail(email)
}

func (s *StaffService) Create(staff *models.User) (int, error) {
	hashedPwd, err := util.HashPassword(*staff.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}
	staff.Password = &hashedPwd
	return s.repo.Create(staff)
}

func (s *StaffService) UpdateFullStaff(user *models.User) (int, error) {
	hashedPwd, err := util.HashPassword(*user.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = &hashedPwd
	return s.repo.UpdateFullStaff(user)
}

func (s *StaffService) Delete(id int) (int, error) {
	return s.repo.Delete(id)
}
