package setting_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsEmailInfoUpdateView	修改邮箱信息api
// @Tags 系统设置
// @summary 修改邮箱信息api
// @Description 修改邮箱信息api
// @Param email body config.Email true "要更新邮箱信息参数"
// @Router /settings_email [put]
// @Produce json
// @success 200 {object} res.Response
func (SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	//绑定json参数
	var email config.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		global.Log.Errorln(err)
		//若是绑定失败 则返回失败信息
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//绑定成功
	//	改变与yaml文件绑定结构体
	global.Config.Email = email
	//	更改配置文件
	err = core.SetYaml()
	if err != nil {
		global.Log.Errorln(err)
		res.FailWithMassage(err.Error(), c)
	}
	res.OkWithMassage("更新成功", c)
	global.Log.Infoln("更改系统信息成功")

}
