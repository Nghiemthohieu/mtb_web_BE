package controllers

import (
	"mtb_web/internal/services"
	"mtb_web/pkg/response"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotificationService *services.NotificationService
}

func NewNotificationController() *NotificationController {
	return &NotificationController{
		NotificationService: services.NewNotificationService(),
	}
}

func (c *NotificationController) GetNotificationsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("userID").(uint)
		notifications, err := c.NotificationService.GetNotificationsByUserID(userID)
		if err != nil {
			response.ErrorRespone(ctx, 500, 20162, err)
			return
		}
		response.SuccessResponse(ctx, 200, notifications)
	}
}

func (c *NotificationController) GetNotificationsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		notifications, err := c.NotificationService.GetNotificationsAdmin()
		if err != nil {
			response.ErrorRespone(ctx, 500, 20163, err)
			return
		}
		response.SuccessResponse(ctx, 200, notifications)
	}
}
