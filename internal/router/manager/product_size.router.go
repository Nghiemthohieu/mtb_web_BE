package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductSizeRouter struct{}

func (p *ProductSizeRouter) InitProductSizeRouter(Router *gin.RouterGroup) {
	ProductSizeRouterPrivate := Router.Group("/admin/product_size")
	{
		ProductSizeRouterPrivate.GET("/list", controllers.NewProductSizeController().GetAll())
		ProductSizeRouterPrivate.POST("/add", controllers.NewProductSizeController().Create())
		ProductSizeRouterPrivate.PUT("/update", controllers.NewProductSizeController().Update())
		ProductSizeRouterPrivate.DELETE("/delete/:id", controllers.NewProductSizeController().Delete())
		ProductSizeRouterPrivate.GET("/get/:id", controllers.NewProductSizeController().GetByID())
		ProductSizeRouterPrivate.POST("/deletes", controllers.NewProductSizeController().DeleteByIDs())
		ProductSizeRouterPrivate.DELETE("/delete", controllers.NewProductSizeController().DeleteAll())
	}
}
