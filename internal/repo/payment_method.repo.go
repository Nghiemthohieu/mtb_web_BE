package repo

import (
	"errors"
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type PaymentMethodRepo struct{}

func NewPaymentMethodRepo() *PaymentMethodRepo {
	return &PaymentMethodRepo{}
}

func (r *PaymentMethodRepo) GetAll() ([]models.PaymentMethod, int, error) {
	var methods []models.PaymentMethod
	if err := global.Mdb.Find(&methods).Error; err != nil {
		return nil, 20138, err
	}
	return methods, 20001, nil
}

func (r *PaymentMethodRepo) GetById(id int) (models.PaymentMethod, int, error) {
	var method models.PaymentMethod
	if err := global.Mdb.First(&method, id).Error; err != nil {
		return models.PaymentMethod{}, 20016, err
	}
	return method, 20137, nil
}

func (r *PaymentMethodRepo) Create(method *models.PaymentMethod) (int, error) {
	var existing models.PaymentMethod
	err := global.Mdb.Unscoped().Where("name = ? OR code = ?", method.Name, method.Code).First(&existing).Error
	if err == nil {
		if existing.DeletedAt.Valid {
			existing.DeletedAt = gorm.DeletedAt{}
			err := global.Mdb.Unscoped().Save(&existing).Error
			if err != nil {
				return 20132, err
			}
			return 20001, nil
		} else {
			return 20132, errors.New("payment method already exists")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := global.Mdb.Create(method).Error; err != nil {
			return 20132, err
		}
		return 20001, nil
	} else {
		return 20132, err
	}
}

func (r *PaymentMethodRepo) Update(method *models.PaymentMethod) (int, error) {
	var existing models.PaymentMethod
	if err := global.Mdb.First(&existing, method.ID).Error; err != nil {
		return 20133, err
	}
	method.CreatedAt = existing.CreatedAt
	if err := global.Mdb.Save(method).Error; err != nil {
		return 20133, err
	}
	return 20001, nil
}

func (r *PaymentMethodRepo) DeleteById(id int) (int, error) {
	var method models.PaymentMethod
	if err := global.Mdb.First(&method, id).Error; err != nil {
		return 20134, err
	}
	if err := global.Mdb.Delete(&method).Error; err != nil {
		return 20134, err
	}
	return 20001, nil
}

func (r *PaymentMethodRepo) DeleteByIDs(ids []int) (int, error) {
	var methods []models.PaymentMethod
	if err := global.Mdb.Where("id IN ?", ids).Find(&methods).Error; err != nil {
		return 2025, err
	}
	if err := global.Mdb.Delete(&methods).Error; err != nil {
		return 20135, err
	}
	return 20001, nil
}

func (r *PaymentMethodRepo) DeleteAll() (int, error) {
	if err := global.Mdb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.PaymentMethod{}).Error; err != nil {
		return 20136, err
	}
	return 20001, nil
}
