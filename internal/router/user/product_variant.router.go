package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductVariantRouter struct {
}

func (p *ProductVariantRouter) InitProductVariantRouter(Router *gin.RouterGroup) {
	productVariantRouterPublic := Router.Group("product_Variant")
	{
		productVariantRouterPublic.GET("list", controllers.NewProductVariantsController().GetAll())
	}
}
