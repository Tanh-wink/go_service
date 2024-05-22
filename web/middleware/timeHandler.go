package middleware

import (
	"fmt"
	"time"
	"strings"
	"github.com/gin-gonic/gin"
	"go/sevice/utils"
	"go/sevice/config"
)

var (
	conf = config.GetConfig()
	logger = utils.GetLogger("Package-middleware")
)

//统计耗时中间件
func HttpTimeHandler(ctx *gin.Context){
	//计时
	start:=time.Now()
	ctx.Next() // 调用后序的处理函数indexHandler
	duration := time.Since(start)
	duration_ms := float64(duration.Nanoseconds())/1000000
	logger.Info(fmt.Sprintf("Total process time:%0.3f ms\n", duration_ms))
	if int(duration_ms) > conf.TimeoutThreshold {
		logger.Warn(fmt.Sprintf("Request Timeout: %0.3f ms!", duration_ms))
		if utils.ContainsString(strings.Split(conf.BotWarningEnv, "/"), conf.Env) {
			err := utils.Bot.SendText(fmt.Sprintf("Servive: %s\nEnv: %s\nRequest Timeout: %0.3f ms!", conf.AppName, conf.Env, duration_ms))
			if err != nil {
				logger.Error("Error sending text message:", err)
			}
		}
	}
}


