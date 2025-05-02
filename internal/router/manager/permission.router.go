package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

func (p *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	PermissionRouterPrivate := Router.Group("admin/permission")
	{
		PermissionRouterPrivate.GET("/list", controllers.NewPermissionController().GetAll())
		PermissionRouterPrivate.POST("/create", controllers.NewPermissionController().Create())
		PermissionRouterPrivate.GET("/details/:id", controllers.NewPermissionController().GetByID())
		PermissionRouterPrivate.PUT("/update", controllers.NewPermissionController().Update())
		PermissionRouterPrivate.DELETE("/delete/:id", controllers.NewPermissionController().Delete())
		PermissionRouterPrivate.POST("/delete", controllers.NewPermissionController().DeleteByIDs())
		PermissionRouterPrivate.DELETE("/delete", controllers.NewPermissionController().DeleteAll())
	}
}
