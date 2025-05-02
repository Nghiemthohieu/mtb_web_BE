package services

import (
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
)

type OrderService struct {
	OrderRepo *repo.OrderRepo
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderRepo: repo.NewOrderRepo(),
	}
}

func (s *OrderService) GetAll() ([]models.Order, int, error) {
	return s.OrderRepo.GetAll()
}

func (s *OrderService) GetByID(id int) (models.Order, int, error) {
	return s.OrderRepo.GetById(id)
}

func (s *OrderService) Create(order *models.Order) (int, error) {
	return s.OrderRepo.Create(order)
}

func (s *OrderService) Update(order *models.Order) (int, error) {
	return s.OrderRepo.Update(order)
}

func (s *OrderService) DeleteById(id int) (int, error) {
	return s.OrderRepo.DeleteById(id)
}

func (s *OrderService) DeleteByIDs(ids []int) (int, error) {
	return s.OrderRepo.DeleteByIDs(ids)
}

func (s *OrderService) DeleteAll() (int, error) {
	return s.OrderRepo.DeleteAll()
}

func (s *OrderService) GetByCustomerID(id int) ([]models.Order, int, error) {
	return s.OrderRepo.GetByCustomerID(id)
}
