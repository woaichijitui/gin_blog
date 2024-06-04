package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingRouter(router *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/setting", settingsApi.SettingInfoView)
}
