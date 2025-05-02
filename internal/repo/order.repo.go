package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type OrderRepo struct{}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

func (r *OrderRepo) GetAll() ([]models.Order, int, error) {
	var orders []models.Order
	if err := global.Mdb.Preload("PaymentMethod").Preload("Status").Find(&orders).Error; err != nil {
		return nil, 20098, err
	}
	return orders, 20001, nil
}

func (r *OrderRepo) GetById(id int) (models.Order, int, error) {
	var order models.Order
	if err := global.Mdb.Preload("PaymentMethod").Preload("Status").First(&order, id).Error; err != nil {
		return models.Order{}, 20016, err
	}
	return order, 20097, nil
}

func (r *OrderRepo) Create(order *models.Order) (int, error) {
	if err := global.Mdb.Create(order).Error; err != nil {
		return 20092, err
	}
	for _, item := range order.ProductOrders {
		var productVariant models.ProductVariant
		if err := global.Mdb.Where("id = ?", item.ProductID).First(productVariant).Error; err == nil {
			return 20092, err
		}
		if productVariant.Stock < item.Quantity {
			return 20059, nil
		}
		productVariant.Stock -= item.Quantity
		if err := global.Mdb.Save(productVariant).Error; err != nil {
			return 20092, err
		}
	}
	return 20001, nil
}

func (r *OrderRepo) Update(order *models.Order) (int, error) {
	var existing models.Order
	if err := global.Mdb.First(&existing, order.ID).Error; err != nil {
		return 20093, err
	}
	order.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(order).Error; err != nil {
		return 20093, err
	}
	return 20001, nil
}

func (r *OrderRepo) DeleteById(id int) (int, error) {
	var order models.Order
	if err := global.Mdb.First(&order, id).Error; err != nil {
		return 20095, err
	}
	if err := global.Mdb.Delete(&order).Error; err != nil {
		return 20094, err
	}
	return 20001, nil
}

func (r *OrderRepo) DeleteByIDs(ids []int) (int, error) {
	var orders []models.Order
	if err := global.Mdb.Where("id IN ?", ids).Find(&orders).Error; err != nil {
		return 2025, err
	}
	if err := global.Mdb.Delete(&orders).Error; err != nil {
		return 20095, err
	}
	return 20001, nil
}

func (r *OrderRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Order{}).Error; err != nil {
		return 20096, err
	}
	return 20001, nil
}

func (r *OrderRepo) GetByCustomerID(id int) ([]models.Order, int, error) {
	var orders []models.Order
	if err := global.Mdb.Preload("PaymentMethod").Preload("ProductOrders").Preload("Status").Where("user_id = ?", id).Find(&orders).Error; err != nil {
		return nil, 20098, err
	}
	return orders, 20001, nil
}
