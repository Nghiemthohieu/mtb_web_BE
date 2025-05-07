package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (p *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/list", controllers.NewProductsController().GetAll())
		productRouterPublic.GET("/get/:id", controllers.NewProductsController().GetByID())
		productRouterPublic.GET("/get_by_category/:id", controllers.NewProductsController().GetByCategoryID())
		productRouterPublic.GET("/get_by_category_and_style/:category_id/:style_id", controllers.NewProductsController().GetByCaregoryIDAndStyleID())
		productRouterPublic.GET("/get_by_style/:style_id", controllers.NewProductsController().GetByStyleID())
		productRouterPublic.GET("/get_by_category_and_color/:category_id/:color_id", controllers.NewProductsController().GetByCaregoryIDAndColorID())
		productRouterPublic.POST("/get_by_category_and_size", controllers.NewProductsController().GetByCategoryIDAndSizeIDs())
		productRouterPublic.GET("/get_by_category_and_material/:category_id/:material_id", controllers.NewProductsController().GetByCaregoryIDAndMaterialID())
	}
}
