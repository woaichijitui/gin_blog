package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

func InitRouter() *gin.Engine {
	//设置gin的模式
	gin.SetMode(global.Config.System.Env)

	router := gin.Default()
	apiGroup := router.Group("/api")
	//系统设置路由
	SettingRouter(apiGroup)
	ImagesRouter(apiGroup)
	AdvertRouter(apiGroup)
	return router
}
