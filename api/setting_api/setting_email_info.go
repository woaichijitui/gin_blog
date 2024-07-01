package setting_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsEmailInfoView 邮箱
// @Tags 系统设置
// @summary 邮箱
// @Description 邮箱
// @Router /settings_email [get]
// @Produce json
// @success 200 {object} res.Response{data=config.Email}
func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {

	res.OkWithData(global.Config.Email, c)

}
