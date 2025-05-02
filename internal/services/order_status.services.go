package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type OrderStatusService struct {
	OrderStatusRepo *repo.OrderStatusRepo
}

func NewOrderStatusService() *OrderStatusService {
	return &OrderStatusService{
		OrderStatusRepo: repo.NewOrderStatusRepo(),
	}
}

func (s *OrderStatusService) GetAll() ([]models.OrderStatus, int, error) {
	return s.OrderStatusRepo.GetAll()
}

func (s *OrderStatusService) GetByID(id int) (models.OrderStatus, int, error) {
	return s.OrderStatusRepo.GetById(id)
}

func (s *OrderStatusService) Create(status *models.OrderStatus) (int, error) {
	return s.OrderStatusRepo.Create(status)
}

func (s *OrderStatusService) Update(status *models.OrderStatus) (int, error) {
	return s.OrderStatusRepo.Update(status)
}

func (s *OrderStatusService) DeleteById(id int) (int, error) {
	return s.OrderStatusRepo.DeleteById(id)
}

func (s *OrderStatusService) DeleteByIDs(ids []int) (int, error) {
	return s.OrderStatusRepo.DeleteByIDs(ids)
}

func (s *OrderStatusService) DeleteAll() (int, error) {
	return s.OrderStatusRepo.DeleteAll()
}
