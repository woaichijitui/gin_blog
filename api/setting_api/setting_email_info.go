package setting_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {

	res.OkWithData(global.Config.Email, c)

}
