package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type PaymentMethodService struct {
	PaymentMethodRepo *repo.PaymentMethodRepo
}

func NewPaymentMethodService() *PaymentMethodService {
	return &PaymentMethodService{
		PaymentMethodRepo: repo.NewPaymentMethodRepo(),
	}
}

func (s *PaymentMethodService) GetAll() ([]models.PaymentMethod, int, error) {
	return s.PaymentMethodRepo.GetAll()
}

func (s *PaymentMethodService) GetByID(id int) (models.PaymentMethod, int, error) {
	return s.PaymentMethodRepo.GetById(id)
}

func (s *PaymentMethodService) Create(method *models.PaymentMethod) (int, error) {
	return s.PaymentMethodRepo.Create(method)
}

func (s *PaymentMethodService) Update(method *models.PaymentMethod) (int, error) {
	return s.PaymentMethodRepo.Update(method)
}

func (s *PaymentMethodService) DeleteById(id int) (int, error) {
	return s.PaymentMethodRepo.DeleteById(id)
}

func (s *PaymentMethodService) DeleteByIDs(ids []int) (int, error) {
	return s.PaymentMethodRepo.DeleteByIDs(ids)
}

func (s *PaymentMethodService) DeleteAll() (int, error) {
	return s.PaymentMethodRepo.DeleteAll()
}
