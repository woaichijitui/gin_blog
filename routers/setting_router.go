package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingRouter(router *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings", settingsApi.SettingsInfoView)
	router.PUT("/settings", settingsApi.SettingsInfoUpdateView)
	router.GET("/settings_email", settingsApi.SettingsEmailInfoView)
	router.PUT("/settings_email", settingsApi.SettingsEmailInfoUpdateView)
}
