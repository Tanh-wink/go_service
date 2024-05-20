package controller

import (
	"github.com/gin-gonic/gin"
	"go/sevice/config"
	"go/sevice/utils"
)

var(
	conf = config.GetConfig()
	logger = utils.GetLogger("controller")
)

func GetProductRecordDetailSC20(ctx *gin.Context){
	logger.Info("GetProductRecordDetailSC20")
	logger.Info(conf.AppName)
}