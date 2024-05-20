package routers

import (
	"encoding/json"
	"go/sevice/utils"
	"go/sevice/web/controller"
	"go/sevice/web/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine{
	// 初始化路由
	router := initRouter()

	v1 := router.Group("/v1")
	{
		v1.GET("/smartPush20/product/v1/getProductRecDetail", controller.GetProductRecordDetailSC20)
		v1.GET("/info", func(ctx *gin.Context) {
			// 异常处理
			// user,err := db.xxx
			// if err != nil {
			// 	panic(utils.ServerError{Msg:"账号或密码错误"})
			// }			
			ctx.JSON(200, gin.H{"msg": "服务启动成功"})
		})
		//接收数据
		v1.POST("/api/get", func(ctx *gin.Context) {
			user := ctx.Param("user")
			ctx.JSON(200, gin.H{"user": user})
		})

		//接收json数据
		v1.POST("/api/getjson", func(ctx *gin.Context) {
			data, _ := ctx.GetRawData()
			var dict map[string]interface{}
			_ = json.Unmarshal(data, &dict)
			ctx.JSON(200, dict)
		})
	}
	return router
}

// 初始化路由
func initRouter() *gin.Engine{
	router := gin.New()
	//add middleware
	router.Use(middleware.HttpTimeHandler, middleware.ExceptionHandler)
	// 设置信任网络 
	// nil 为不计算，避免性能消耗，上线应当设置router.SetTrustedProxies([]string{"10.0.x.1", "10.0.0.x", "10.0.0.x"}
	_ = router.SetTrustedProxies(nil)
	
	//Load home
	router.LoadHTMLGlob(utils.ProjectPath + "/web/template/static/*")
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"msg": "This is a go web project",
		})
	})

	return router
}