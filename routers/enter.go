package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
	"gvb_server/utils/common"
)

func InitRouter() *gin.Engine {
	//设置gin的模式
	gin.SetMode(global.Config.System.Env)

	router := gin.Default()
	//
	router.GET("swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("required_bool", common.RequiredBool)
	}

	apiGroup := router.Group("/api")

	//系统设置路由
	SettingRouter(apiGroup)
	ImagesRouter(apiGroup)
	AdvertRouter(apiGroup)
	MenuRouter(apiGroup)
	return router
}
