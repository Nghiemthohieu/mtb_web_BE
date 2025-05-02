package repo

import (
	"errors"
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type OrderStatusRepo struct{}

func NewOrderStatusRepo() *OrderStatusRepo {
	return &OrderStatusRepo{}
}

func (r *OrderStatusRepo) GetAll() ([]models.OrderStatus, int, error) {
	var statuses []models.OrderStatus
	if err := global.Mdb.Find(&statuses).Error; err != nil {
		return nil, 20088, err
	}
	return statuses, 20001, nil
}

func (r *OrderStatusRepo) GetById(id int) (models.OrderStatus, int, error) {
	var status models.OrderStatus
	if err := global.Mdb.First(&status, id).Error; err != nil {
		return models.OrderStatus{}, 20016, err
	}
	return status, 20087, nil
}

func (r *OrderStatusRepo) Create(status *models.OrderStatus) (int, error) {
	var existing models.OrderStatus
	err := global.Mdb.Unscoped().Where("status = ?", status.Status).First(&existing).Error
	if err == nil {
		if existing.DeletedAt.Valid {
			existing.DeletedAt = gorm.DeletedAt{}
			err := global.Mdb.Unscoped().Save(&existing).Error
			if err != nil {
				return 20082, err
			}
			return 20001, nil
		} else {
			return 20082, errors.New("order status already exists")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := global.Mdb.Create(status).Error; err != nil {
			return 20082, err
		}
		return 20001, nil
	} else {
		return 20082, err
	}
}

func (r *OrderStatusRepo) Update(status *models.OrderStatus) (int, error) {
	var existing models.OrderStatus
	if err := global.Mdb.First(&existing, status.ID).Error; err != nil {
		return 20083, err
	}
	status.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(status).Error; err != nil {
		return 20083, err
	}
	return 20001, nil
}

func (r *OrderStatusRepo) DeleteById(id int) (int, error) {
	var status models.OrderStatus
	if err := global.Mdb.First(&status, id).Error; err != nil {
		return 20084, err
	}
	if err := global.Mdb.Delete(&status).Error; err != nil {
		return 20084, err
	}
	return 20001, nil
}

func (r *OrderStatusRepo) DeleteByIDs(ids []int) (int, error) {
	var statuses []models.OrderStatus
	if err := global.Mdb.Where("id IN ?", ids).Find(&statuses).Error; err != nil {
		return 2025, err
	}
	if err := global.Mdb.Delete(&statuses).Error; err != nil {
		return 20085, err
	}
	return 20001, nil
}

func (r *OrderStatusRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.OrderStatus{}).Error; err != nil {
		return 20086, err
	}
	return 20001, nil
}
