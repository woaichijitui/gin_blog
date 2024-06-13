package setting_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsInfoUpdateView	修改配置文件信息api
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	//绑定json参数
	var siteInfo config.SiteInfo
	err := c.ShouldBindJSON(&siteInfo)
	if err != nil {
		global.Log.Errorln(err)
		//若是绑定失败 则返回失败信息
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//绑定成功
	//	改变与yaml文件绑定结构体
	global.Config.SiteInfo = siteInfo
	//	更改配置文件
	err = core.SetYaml()
	if err != nil {
		global.Log.Errorln(err)
		res.FailWithMassage(err.Error(), c)
	}
	res.OkWithMassage("更新成功", c)
	global.Log.Infoln("更改系统信息成功")

}
