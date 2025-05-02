package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductStyleRouter struct{}

func (p *ProductStyleRouter) InitProductStyleRouter(Router *gin.RouterGroup) {
	productStyleRouterPrivate := Router.Group("admin/product_style")
	{
		productStyleRouterPrivate.GET("/list", controllers.NewProductStylesController().GetAll())
		productStyleRouterPrivate.POST("/create", controllers.NewProductStylesController().Create())
		productStyleRouterPrivate.GET("/details/:id", controllers.NewProductStylesController().GetByID())
		productStyleRouterPrivate.PUT("/update", controllers.NewProductStylesController().Update())
		productStyleRouterPrivate.DELETE("/delete/:id", controllers.NewProductStylesController().Delete())
		productStyleRouterPrivate.POST("/delete", controllers.NewProductStylesController().DeleteByIDs())
		productStyleRouterPrivate.DELETE("/delete", controllers.NewProductStylesController().DeleteAll())
	}
}
