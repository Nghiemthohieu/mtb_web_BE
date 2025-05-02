package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductStyleRouter struct {
}

func (p *ProductStyleRouter) InitProductStyleRouter(Router *gin.RouterGroup) {
	productStyleRouterPublic := Router.Group("product_style")
	{
		productStyleRouterPublic.GET("list", controllers.NewProductStylesController().GetAll())
	}
}
