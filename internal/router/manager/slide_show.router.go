package manager

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type SlideShowRouter struct{}

func (p *SlideShowRouter) InitSlideShowRouter(Router *gin.RouterGroup) {
	SlideShowRouterPrivate := Router.Group("admin/slide_show")
	{
		SlideShowRouterPrivate.GET("/list", controllers.NewSlideShowController().GetAll())
		SlideShowRouterPrivate.POST("/create", controllers.NewSlideShowController().Create())
		SlideShowRouterPrivate.GET("/details/:id", controllers.NewSlideShowController().GetByID())
		SlideShowRouterPrivate.PUT("/update", controllers.NewSlideShowController().Update())
		SlideShowRouterPrivate.DELETE("/delete/:id", controllers.NewSlideShowController().Delete())
	}
}
