package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (p *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouterPrivate := Router.Group("admin/product")
	{
		productRouterPrivate.GET("/list", controllers.NewProductsController().GetAll())
		productRouterPrivate.POST("/create", controllers.NewProductsController().Create())
		productRouterPrivate.GET("/details/:id", controllers.NewProductsController().GetByID())
		productRouterPrivate.PUT("/update", controllers.NewProductsController().Update())
		productRouterPrivate.DELETE("/delete/:id", controllers.NewProductsController().Delete())
		productRouterPrivate.POST("/delete", controllers.NewProductsController().DeleteByIDs())
		productRouterPrivate.DELETE("/delete", controllers.NewProductsController().DeleteAll())
	}
}
