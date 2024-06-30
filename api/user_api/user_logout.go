package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
)

// UserLogoutView 用户注销
// @Tags 用户管理
// @summary 用户注销
// @Description 用户注销
// @Param token header string  true "token"
// @Router /user_logout [get]
// @Produce json
// @success 200 {object} res.Response、
func (UserApi) UserLogoutView(c *gin.Context) {

	token := c.GetHeader("token")

	exp := service.Service.UserService.GetTokenExp(c)

	//	 将token和过期存入redis
	err := service.Service.RedisService.SetLogoutToken(token, exp)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("注销失败", c)
		return
	}

	//	成功响应
	res.OkWithMassage("注销成功", c)

}
