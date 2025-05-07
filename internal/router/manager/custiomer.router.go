package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (p *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	CustomerRouterPrivate := Router.Group("admin/customer")
	{
		CustomerRouterPrivate.GET("/list", controllers.NewCustomerController().GetAll())
		CustomerRouterPrivate.GET("/details/:id", controllers.NewCustomerController().GetByID())
		CustomerRouterPrivate.GET("/details_email/:email", controllers.NewCustomerController().GetByEmail())
		CustomerRouterPrivate.DELETE("/delete/:id", controllers.NewCustomerController().Delete())
	}
}
