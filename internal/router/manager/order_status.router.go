package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type OrderStatusRouter struct{}

func (p *OrderStatusRouter) InitOrderStatusRouter(Router *gin.RouterGroup) {
	OrderStatusRouterPrivate := Router.Group("admin/order_status")
	{
		OrderStatusRouterPrivate.GET("/list", controllers.NewOrderStatusController().GetAll())
		OrderStatusRouterPrivate.POST("/create", controllers.NewOrderStatusController().Create())
		OrderStatusRouterPrivate.GET("/details/:id", controllers.NewOrderStatusController().GetByID())
		OrderStatusRouterPrivate.PUT("/update", controllers.NewOrderStatusController().Update())
		OrderStatusRouterPrivate.DELETE("/delete/:id", controllers.NewOrderStatusController().Delete())
		OrderStatusRouterPrivate.POST("/delete", controllers.NewOrderStatusController().DeleteByIDs())
		OrderStatusRouterPrivate.DELETE("/delete", controllers.NewOrderStatusController().DeleteAll())
	}
}
