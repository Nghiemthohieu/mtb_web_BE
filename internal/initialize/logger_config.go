package initialize

import (
	"mtb_web/global"
	"mtb_web/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
