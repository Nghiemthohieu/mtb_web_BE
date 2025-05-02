package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductColorRouter struct {
}

func (p *ProductColorRouter) InitProductColorRouter(Router *gin.RouterGroup) {
	productColorRouterPublic := Router.Group("product_color")
	{
		productColorRouterPublic.GET("list", controllers.NewProductColorController().GetAll())
	}
}
