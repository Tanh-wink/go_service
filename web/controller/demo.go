package controller

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "go/sevice/config"
	"go/sevice/dto"
	"go/sevice/utils"
	"net/http"
	"go/sevice/web/middleware"
)

var (
	// conf = config.GetConfig()
	logger = utils.GetLogger("controller")
)

func GetProductRecordDetailSC20(ctx *gin.Context) {
	logger.Info("GetProductRecordDetailSC20")
	// 请求参数接收与验证
	var request dto.ProductRecRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		logger.Error(err.Error())
		panic(middleware.ParamError{Msg: err.Error()})
	}
	
	// 业务逻辑代码
	var products = []string{"product1", "product2", "product3"}
	productRecResponse := &dto.ProductRecResponse{}
	productRecResponse.Hot = products

	// 返回响应结果
	response := dto.NewResponse(http.StatusOK, "success", productRecResponse)
	ctx.JSON(http.StatusOK, response)
}
