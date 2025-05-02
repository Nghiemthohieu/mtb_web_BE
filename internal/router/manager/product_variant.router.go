package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductVariantRouter struct{}

func (p *ProductVariantRouter) InitProductVariantRouter(Router *gin.RouterGroup) {
	productVariantRouterPrivate := Router.Group("admin/product_Variant")
	{
		productVariantRouterPrivate.GET("/list", controllers.NewProductVariantsController().GetAll())
		productVariantRouterPrivate.POST("/create", controllers.NewProductVariantsController().Create())
		productVariantRouterPrivate.GET("/details/:id", controllers.NewProductVariantsController().GetByID())
		productVariantRouterPrivate.PUT("/update", controllers.NewProductVariantsController().Update())
		productVariantRouterPrivate.DELETE("/delete/:id", controllers.NewProductVariantsController().Delete())
		productVariantRouterPrivate.POST("/delete", controllers.NewProductVariantsController().DeleteByIDs())
		productVariantRouterPrivate.DELETE("/delete", controllers.NewProductVariantsController().DeleteAll())
	}
}
