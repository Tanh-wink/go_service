package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/sevice/config"
	"go/sevice/utils"
	"go/sevice/web/routers"
)

var (
	conf   = config.GetConfig()
	logger = utils.GetLogger("Package-main")
)

func main() {
	// 初始化环境
	initEnv()

	server := routers.InitRouter()
	err := server.Run(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		logger.Info("Service startup failed!")
	}
}

func initEnv() {
	// 自动识别环境
	if utils.IsServer() {
		// 服务器环境下的配置
		gin.SetMode(gin.ReleaseMode)   // Release 模式：在生产环境中使用
		logger.Info(fmt.Sprintf("Current App is running on :%s", utils.KIND_SERVER))
	} else {
		gin.SetMode(gin.DebugMode)     // Debug 模式: 在开发和调试阶段使用
		logger.Info(fmt.Sprintf("Current App is running on :%s", utils.KIND_LOCAL))
	}
	logger.Info(fmt.Sprintf("Current Environment is :%s", conf.Env))
}
