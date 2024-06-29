package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/service_com"
	"gvb_server/utils"
)

// UserListView 用户list
// @Tags 用户管理
// @summary 用户list
// @Description 用户list
// @Param page query models.PageInfo true "表示单个参数"
// @Router /users [get]
// @Produce json
// @success 200 {object} res.Response{data=[]models.UserModel}
func (UserApi) UserListView(c *gin.Context) {
	//判断是否是管理者
	token := c.GetHeader("token")
	//有无token
	if token == "" {
		res.FailWithMassage("无token", c)
		return
	}
	//解析token
	claims, err := utils.ParseTokenRs256(token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("token解析失败", c)
		return
	}
	if ctype.Role(claims.Role) != ctype.PermissionAdmin {
		//	若不是管理员
		res.FailWithMassage("非管理员用户", c)
		return
	}

	//	绑定参数
	var page models.PageInfo
	//绑定参数
	err = c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userModel models.UserModel
	//	用户列表
	list, count, err := service_com.ComList(userModel, service_com.Option{page, true})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage(err.Error(), c)
		return
	}

	//管理员用户不显示username
	var userList []models.UserModel
	for _, user := range list {
		if user.Role == ctype.PermissionAdmin {
			user.UserName = ""
		}
		userList = append(userList, user)
	}

	//电话邮箱脱敏处理
	//	响应
	res.OkWithList(userList, count, c)
}
