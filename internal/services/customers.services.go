package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type CustomerService struct {
	CustomerRepo *repo.CustomerRepo
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		CustomerRepo: repo.NewCustomerRepo(),
	}
}

func (s *CustomerService) GetAll() ([]models.User, int, error) {
	return s.CustomerRepo.GetAll()
}

func (s *CustomerService) GetByID(id int) (models.User, int, error) {
	return s.CustomerRepo.GetByID(id)
}

func (s *CustomerService) GetByEmail(email string) (models.User, int, error) {
	return s.CustomerRepo.GetByEmail(email)
}

func (s *CustomerService) Delete(id int) (int, error) {
	return s.CustomerRepo.Delete(id)
}
