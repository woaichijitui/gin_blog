package setting_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {

	res.Ok(map[string]string{"name": "htt"}, "成功", c)

}
