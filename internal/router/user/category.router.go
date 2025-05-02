package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (c *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouterPublic := Router.Group("/category")
	{
		CategoryRouterPublic.GET("/list", controllers.NewCategoriesController().GetAll())
		CategoryRouterPublic.GET("/get/:id", controllers.NewCategoriesController().GetByID())
	}
}
