package middleware

import (
	"fmt"
	"go/sevice/utils"
	"net/http"
	"runtime/debug"
	"github.com/gin-gonic/gin"
)


func ExceptionHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
     		// 修改错误信息解析
			errCode, errMsg := getError(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg": errMsg,
			})
			ctx.Abort()
		}
	}()
	ctx.Next()
}

var errorResponse = utils.NewErrorResponse()

func getError(err interface{}) (int, string) {
	switch v := err.(type) {
		
		case utils.ParamError:
			// 符合预期的错误，可以直接返回给客户端
			return errorResponse.Error_Param.Code, v.Msg
		case utils.ServerError:
			// 符合预期的错误，可以直接返回给客户端
			return errorResponse.Error_Server.Code, v.Msg
		case utils.LoadError:
			// 符合预期的错误，可以直接返回给客户端
			return errorResponse.Error_Load.Code, v.Msg
		case utils.InferError:
			// 符合预期的错误，可以直接返回给客户端
			return errorResponse.Error_Infer.Code, v.Msg
		case error:
			// 一律返回服务器错误，避免返回堆栈错误给客户端，实际还可以针对系统错误做其他处理
			debug.PrintStack()
			fmt.Printf("panic: %v\n", v.Error())
			return 1000, "Service error!"
		default:
		// 同上
			return 1000, "Unkown error!"
	}
}