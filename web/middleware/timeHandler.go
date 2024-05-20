package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//统计耗时中间件
func HttpTimeHandler(ctx *gin.Context){
	//计时
	start:=time.Now()
	ctx.Next() // 调用后序的处理函数indexHandler
	cost := time.Since(start)
	fmt.Printf("cost:%0.3f ms\n", float64(cost.Nanoseconds())/1000000)
}


