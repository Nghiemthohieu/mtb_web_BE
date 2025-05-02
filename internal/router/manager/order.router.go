package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (p *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouterPrivate := Router.Group("admin/order")
	{
		OrderRouterPrivate.GET("/list", controllers.NewOrderController().GetAll())
		OrderRouterPrivate.POST("/create", controllers.NewOrderController().Create())
		OrderRouterPrivate.GET("/details/:id", controllers.NewOrderController().GetByID())
		OrderRouterPrivate.PUT("/update", controllers.NewOrderController().Update())
		OrderRouterPrivate.DELETE("/delete/:id", controllers.NewOrderController().Delete())
		OrderRouterPrivate.POST("/delete", controllers.NewOrderController().DeleteByIDs())
		OrderRouterPrivate.DELETE("/delete", controllers.NewOrderController().DeleteAll())
	}
}
