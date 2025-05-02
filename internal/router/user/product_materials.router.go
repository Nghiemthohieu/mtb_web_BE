package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductMaterialsRouter struct {
}

func (p *ProductMaterialsRouter) InitProductMaterialsRouter(Router *gin.RouterGroup) {
	productMaterialsRouterPublic := Router.Group("/product_materials")
	{
		productMaterialsRouterPublic.GET("/list", controllers.NewProductMaterialsController().GetAll())
	}
}
