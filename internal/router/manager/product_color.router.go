package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ProductColorRouter struct{}

func (p *ProductColorRouter) InitProductColorRouter(Router *gin.RouterGroup) {
	productColorRouter := Router.Group("admin/product_color")
	{
		productColorRouter.POST("/add", controllers.NewProductColorController().Create())              // 添加产品颜色
		productColorRouter.GET("/list", controllers.NewProductColorController().GetAll())              // 获取产品颜色列表
		productColorRouter.GET("/get/:id", controllers.NewProductColorController().GetById())          // 获取单个产品颜色
		productColorRouter.PUT("/update", controllers.NewProductColorController().Update())            // 更新产品颜色
		productColorRouter.DELETE("/delete/:id", controllers.NewProductColorController().DeleteById()) // 删除产品颜色
		productColorRouter.DELETE("/delete", controllers.NewProductColorController().DeleteAll())      // 批量删除产品颜色
		productColorRouter.POST("/delete", controllers.NewProductColorController().DeleteByIDs())      // 批量删除产品颜色
	}
}
