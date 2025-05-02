package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type StaffsRouter struct{}

func (p *StaffsRouter) InitStaffsRouter(Router *gin.RouterGroup) {
	StaffsRouterPrivate := Router.Group("admin/staffs")
	{
		StaffsRouterPrivate.GET("/list", controllers.NewStaffsController().GetAllStaffs())
		StaffsRouterPrivate.POST("/create", controllers.NewStaffsController().Create())
		StaffsRouterPrivate.GET("/details/:id", controllers.NewStaffsController().GetByID())
		StaffsRouterPrivate.PUT("/update", controllers.NewStaffsController().Update())
		StaffsRouterPrivate.DELETE("/delete/:id", controllers.NewStaffsController().Delete())
	}
}
