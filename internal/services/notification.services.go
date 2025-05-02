package services

import (
	"encoding/json"
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	"mtb_web/internal/websocket"
)

type NotificationService struct {
	NotificationRepo *repo.NotificationRepo
	Hub              *websocket.Hub
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		NotificationRepo: repo.NewNotificationRepo(),
		Hub:              websocket.NewHub(),
	}
}

func (n *NotificationService) GetNotificationsByUserID(userID uint) ([]models.Notification, error) {
	return n.NotificationRepo.GetNotificationsByUserID(userID)
}

func (n *NotificationService) GetNotificationsAdmin() ([]models.Notification, error) {
	return n.NotificationRepo.GetNotificationsAdmin()
}

func (n *NotificationService) OrderNotification(userID uint, orderID uint) error {
	// Notification cho user
	notif := models.Notification{
		UserID:  &userID, // ✅ fix ở đây
		Title:   "Đơn hàng mới",
		Message: "Bạn vừa đặt hàng thành công!",
		Type:    "order",
		IsRead:  false,
	}
	if err := n.NotificationRepo.CreateNotification(&notif); err != nil {
		return err
	}

	payload := map[string]interface{}{
		"title":   notif.Title,
		"message": notif.Message,
		"orderID": orderID,
	}
	msg, _ := json.Marshal(payload)
	n.Hub.SendToUser(userID, msg)

	// Notification cho admin
	notifAdmin := models.Notification{
		UserID:  nil, // ✅ OK
		Title:   "Đơn hàng mới",
		Message: fmt.Sprintf("User %d đã đặt đơn hàng %d", userID, orderID),
		Type:    "order",
		IsRead:  false,
	}
	if err := n.NotificationRepo.CreateNotification(&notifAdmin); err != nil {
		return err
	}

	payload = map[string]interface{}{
		"title":   notifAdmin.Title,
		"message": notifAdmin.Message,
		"orderID": orderID,
	}
	msg, _ = json.Marshal(payload)

	adminID := uint(1) // Hardcode ID admin
	n.Hub.SendToUser(adminID, msg)

	return nil
}
