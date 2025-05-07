package initialize

import (
	"mtb_web/global"
	"mtb_web/internal/middlewares"
	"mtb_web/internal/router"
	"mtb_web/internal/websocket"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//middleware
	r.Use(middlewares.LoggerMiddle())        //logging
	r.Use(middlewares.CorsMiddleware())      //cross
	r.Use(middlewares.RateLimitMiddleware()) //limiter global
	r.Use(middlewares.EnsureSessionID())
	managerRouter := router.RouterGroupApp.Manager
	userRouter := router.RouterGroupApp.User

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitProductSizeRouter(mainGroup)
		userRouter.InitProductColorRouter(mainGroup)
		userRouter.InitProductMaterialsRouter(mainGroup)
		userRouter.InitProductStyleRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
		userRouter.InitCategoryRouter(mainGroup)
		userRouter.InitProductVariantRouter(mainGroup)
		userRouter.InitSlideShowRouter(mainGroup)
	}
	{
		managerRouter.InitProductSizeRouter(mainGroup)
		managerRouter.InitProductColorRouter(mainGroup)
		managerRouter.InitProductMaterialsRouter(mainGroup)
		managerRouter.InitProductStyleRouter(mainGroup)
		managerRouter.InitProductRouter(mainGroup)
		managerRouter.InitCategoryRouter(mainGroup)
		managerRouter.InitProductVariantRouter(mainGroup)
		managerRouter.InitSlideShowRouter(mainGroup)
		managerRouter.InitStaffsRouter(mainGroup)
		managerRouter.InitCustomerRouter(mainGroup)
	}
	return r
}

func SetupRoutes(r *gin.Engine, hub *websocket.Hub) {
	r.GET("/ws", websocket.ServeWs(hub))
}
