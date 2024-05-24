package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"github.com/gin-gonic/gin"
)

type ErrorCode struct {
	UnkownErrorCode		int
	ParamErrorCode		int
	ServerErrorCode		int
	LoadErrorCode		int
	InferErrorCode		int
	
}

func NewErrorCode() *ErrorCode {
	return &ErrorCode{
		UnkownErrorCode: 	1000,
		ParamErrorCode: 	300,
		ServerErrorCode: 	500,
		LoadErrorCode: 		1001,
		InferErrorCode: 	1002,
	}
}

type ParamError struct {
	Msg 	string `json:"msg"`
}

type LoadError struct {
	Msg 	string `json:"msg"`
}

type InferError struct {
	Msg 	string `json:"msg"`
}

type ServerError struct {
	Msg 	string `json:"msg"`
}

func getError(err interface{}) (int, string) {
	switch v := err.(type) {
		case ParamError:
			// 符合预期的错误，可以直接返回给客户端
			return errorCode.ParamErrorCode, v.Msg
		case LoadError:
			return errorCode.LoadErrorCode, v.Msg
		case InferError:
			return errorCode.InferErrorCode, v.Msg

		case error:
			// 一律返回服务器错误，避免返回堆栈错误给客户端，实际还可以针对系统错误做其他处理
			debug.PrintStack()
			fmt.Printf("panic: %v\n", v.Error())
			return errorCode.ServerErrorCode, "Service error!"
		default:
			return errorCode.UnkownErrorCode, "Unkown error!"
	}
}

var errorCode = NewErrorCode()

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