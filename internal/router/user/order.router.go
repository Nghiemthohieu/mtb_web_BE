package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (c *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouterPublic := Router.Group("/Order")
	{
		OrderRouterPublic.GET("/list", controllers.NewOrderController().GetAll())
		OrderRouterPublic.GET("/get/:id", controllers.NewOrderController().GetByID())
		OrderRouterPublic.POST("/create", controllers.NewOrderController().Create())
		OrderRouterPublic.POST("/delete", controllers.NewOrderController().Delete())
		OrderRouterPublic.GET("/get_by_user_id/:id", controllers.NewOrderController().GetByCustomerID())
	}
}
