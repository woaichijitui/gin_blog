package setting_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsInfoView 系统信息
// @Tags 系统设置
// @summary 系统信息
// @Description 系统信息
// @Router /settings [get]
// @Produce json
// @success 200 {object} res.Response{data=config.SiteInfo}
func (SettingsApi) SettingsInfoView(c *gin.Context) {

	res.OkWithData(global.Config.SiteInfo, c)

}
