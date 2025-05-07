package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"

	"gorm.io/gorm"
)

type StaffRepo struct {
}

func NewStaffRepo() *StaffRepo {
	return &StaffRepo{}
}

func (s *StaffRepo) GetAllStaff() ([]models.User, int, error) {
	var Staffs []models.User
	if err := global.Mdb.Joins("JOIN go_db_staffs ON go_db_staffs.user_id = go_db_users.id").Find(&Staffs).Error; err != nil {
		return nil, 20098, err
	}
	return Staffs, 20001, nil
}

func (s *StaffRepo) GetStaffById(id int) (models.User, int, error) {
	var Staffs models.User
	if err := global.Mdb.Joins("JOIN go_db_staffs ON go_db_staffs.user_id = go_db_users.id").First(&Staffs, id).Error; err != nil {
		return models.User{}, 20016, err
	}
	return Staffs, 20097, nil
}

func (s *StaffRepo) GetStaffByEmail(email string) (models.User, int, error) {
	var Staffs models.User
	if err := global.Mdb.Joins("JOIN go_db_staffs ON go_db_staffs.user_id = go_db_users.id").Where("email = ?", email).First(&Staffs).Error; err != nil {
		return models.User{}, 20016, err
	}
	return Staffs, 20097, nil
}

func (s *StaffRepo) Create(staff *models.User) (int, error) {
	if err := global.Mdb.Create(staff).Error; err != nil {
		return 20098, err
	}
	return 20001, nil
}

func (s *StaffRepo) UpdateFullStaff(user *models.User) (int, error) {
	err := global.Mdb.Transaction(func(tx *gorm.DB) error {
		// Cập nhật user
		if err := tx.Save(user).Error; err != nil {
			return err
		}

		// Cập nhật thông tin staff nếu có
		if user.Staff != nil {
			user.Staff.UserID = user.ID // đảm bảo đúng
			if err := tx.Save(user.Staff).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return 20099, err // hoặc mã lỗi riêng cho update fail
	}

	return 20001, nil // success code
}

func (s *StaffRepo) Delete(id int) (int, error) {
	if err := global.Mdb.Delete(&models.User{}, id).Error; err != nil {
		return 20098, err
	}
	return 20001, nil
}
