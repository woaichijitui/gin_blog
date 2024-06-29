package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
)

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"用户名不存在"`
	Password string `json:"password" binding:"required" msg:"密码不正确"`
}

// EmailLoginView 邮箱或者用户名登录
// @Tags 用户管理
// @summary 邮箱登录
// @Description 邮箱登录
// @Param cr body LoginRequest true "用户 密码 "
// @Router /email_login [post]
// @Produce json
// @success 200 {object} res.Response
func (l UserApi) EmailLoginView(c *gin.Context) {
	//	接收参数
	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//	判断是否有该用户
	var userModel models.UserModel
	row := global.DB.Find(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).RowsAffected
	if row == 0 {
		res.FailWithMassage("用户名错误", c)
		return
	}

	//	密码验证
	ok := utils.PasswordVerify(cr.Password, userModel.Password)
	if !ok {
		res.FailWithMassage("密码输入错误", c)
		return
	}

	//	生成token
	token, err := utils.GenerateTokenUsingRS256(userModel.ID, userModel.UserName, userModel.Role)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("生成token错误", c)
		return
	}
	//	响应
	res.OkWithData(token, c)
}
