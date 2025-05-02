package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (p *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouterPrivate := Router.Group("admin/role")
	{
		RoleRouterPrivate.GET("/list", controllers.NewRoleController().GetAll())
		RoleRouterPrivate.POST("/create", controllers.NewRoleController().Create())
		RoleRouterPrivate.GET("/details/:id", controllers.NewRoleController().GetByID())
		RoleRouterPrivate.PUT("/update", controllers.NewRoleController().Update())
		RoleRouterPrivate.DELETE("/delete/:id", controllers.NewRoleController().Delete())
		RoleRouterPrivate.POST("/delete", controllers.NewRoleController().DeleteByIDs())
		RoleRouterPrivate.DELETE("/delete", controllers.NewRoleController().DeleteAll())
	}
}
