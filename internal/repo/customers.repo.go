package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"
)

type CustomerRepo struct{}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{}
}

func (c *CustomerRepo) GetAll() ([]models.User, int, error) {
	var customers []models.User
	if err := global.Mdb.Joins("JOIN go_db_customers ON go_db_customers.user_id = go_db_users.id").Find(&customers).Error; err != nil {
		return nil, 20078, err
	}
	return customers, 20001, nil
}

func (c *CustomerRepo) GetByID(id int) (models.User, int, error) {
	var customers models.User
	if err := global.Mdb.Joins("JOIN go_db_customers ON go_db_customers.user_id = go_db_users.id").First(&customers, id).Error; err != nil {
		return models.User{}, 20078, err
	}
	return customers, 20001, nil
}

func (c *CustomerRepo) GetByEmail(email string) (models.User, int, error) {
	var customers models.User
	if err := global.Mdb.Joins("JOIN go_db_customers ON go_db_customers.user_id = go_db_users.id").Where("email = ?", email).First(&customers).Error; err != nil {
		return models.User{}, 20078, err
	}
	return customers, 20001, nil
}

func (c *CustomerRepo) Delete(id int) (int, error) {
	if err := global.Mdb.Delete(&models.User{}, id).Error; err != nil {
		return 20078, err
	}
	return 20001, nil
}