package initialize

import (
	"fmt"
	"mtb_web/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration MySql...", global.Config.Mysql.UserName)
	InitLogger()
	global.Logger.Info("Config log ok!!!", zap.String("ok", "Success"))
	InitMySQL()
	InitRedis()

	// Kiểm tra lỗi AWS trước khi tiếp tục
	defer func() {
		if r := recover(); r != nil {
			global.Logger.Error("Server không thể khởi động do lỗi AWS", zap.Any("error", r))
			panic(r)
		}
	}()
	InitAws()

	r := NewRouter()

	// Start server on port 8080
	r.Run()
}
