package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductMaterialsRouter struct{}

func (p *ProductMaterialsRouter) InitProductMaterialsRouter(Router *gin.RouterGroup) {
	productMaterialsRouter := Router.Group("admin/product_materrials")
	{
		productMaterialsRouter.POST("/add", controllers.NewProductMaterialsController().Create())          // 添加产品颜色
		productMaterialsRouter.GET("/list", controllers.NewProductMaterialsController().GetAll())          // 获取产品颜色列表
		productMaterialsRouter.GET("/get/:id", controllers.NewProductMaterialsController().GetByID())      // 获取单个产品颜色
		productMaterialsRouter.PUT("/update", controllers.NewProductMaterialsController().Update())        // 更新产品颜色
		productMaterialsRouter.DELETE("/delete/:id", controllers.NewProductMaterialsController().Delete()) // 删除产品颜色
		productMaterialsRouter.DELETE("/delete", controllers.NewProductMaterialsController().DeleteAll())  // 批量删除产品颜色
		productMaterialsRouter.POST("/delete", controllers.NewProductMaterialsController().DeleteByIDs())  // 批量删除产品颜色
	}
}
