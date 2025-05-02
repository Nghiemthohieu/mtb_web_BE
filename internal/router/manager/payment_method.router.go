package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type PaymentMethodRouter struct{}

func (p *PaymentMethodRouter) InitPaymentMethodRouter(Router *gin.RouterGroup) {
	PaymentMethodRouterPrivate := Router.Group("admin/payment_method")
	{
		PaymentMethodRouterPrivate.GET("/list", controllers.NewPaymentMethodController().GetAll())
		PaymentMethodRouterPrivate.POST("/create", controllers.NewPaymentMethodController().Create())
		PaymentMethodRouterPrivate.GET("/details/:id", controllers.NewPaymentMethodController().GetByID())
		PaymentMethodRouterPrivate.PUT("/update", controllers.NewPaymentMethodController().Update())
		PaymentMethodRouterPrivate.DELETE("/delete/:id", controllers.NewPaymentMethodController().Delete())
		PaymentMethodRouterPrivate.POST("/delete", controllers.NewPaymentMethodController().DeleteByIDs())
		PaymentMethodRouterPrivate.DELETE("/delete", controllers.NewPaymentMethodController().DeleteAll())
	}
}
