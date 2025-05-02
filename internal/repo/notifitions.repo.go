package repo

import (
	"mtb_web/global"
	"mtb_web/internal/models"
)

type NotificationRepo struct {
}

func NewNotificationRepo() *NotificationRepo {
	return &NotificationRepo{}
}

func (n *NotificationRepo) CreateNotification(notification *models.Notification) error {
	if err := global.Mdb.Create(&notification).Error; err != nil {
		return err
	}
	return nil
}

func (n *NotificationRepo) GetNotificationsByUserID(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := global.Mdb.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *NotificationRepo) GetNotificationsAdmin() ([]models.Notification, error) {
	var notifications []models.Notification
	if err := global.Mdb.Where("user_id IS NULL").Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}
