package user

import (
	"mtb_web/internal/controllers"

	"github.com/gin-gonic/gin"
)

type SlideShowRouter struct {
}

func (p *SlideShowRouter) InitSlideShowRouter(Router *gin.RouterGroup) {
	SlideShowRouterPublic := Router.Group("/slide_show")
	{
		SlideShowRouterPublic.GET("/list", controllers.NewSlideShowController().GetAll())
		SlideShowRouterPublic.GET("/get/:id", controllers.NewSlideShowController().GetByID())
	}
}
