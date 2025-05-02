package router

import (
	"mtb_web/internal/router/manager"
	"mtb_web/internal/router/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
