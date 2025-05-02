package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (p *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouterPrivate := Router.Group("admin/category")
	{
		CategoryRouterPrivate.GET("/list", controllers.NewCategoriesController().GetAll())
		CategoryRouterPrivate.POST("/create", controllers.NewCategoriesController().Create())
		CategoryRouterPrivate.GET("/details/:id", controllers.NewCategoriesController().GetByID())
		CategoryRouterPrivate.PUT("/update/:id", controllers.NewCategoriesController().Update())
		CategoryRouterPrivate.DELETE("/delete/:id", controllers.NewCategoriesController().Delete())
		CategoryRouterPrivate.POST("/delete", controllers.NewCategoriesController().DeleteByIDs())
		CategoryRouterPrivate.DELETE("/delete", controllers.NewCategoriesController().DeleteAll())
	}
}
