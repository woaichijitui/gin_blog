package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils"
)

type UserUpdatePwdRequest struct {
	UserID      uint   `json:"user_id" binding:"required" msg:"请输入用户id"`
	Password    string `json:"password" binding:"required" msg:"请输入密码"`
	NewPassword string `json:"new_password" binding:"required" msg:"请输入新密码"`
	RePassword  string `json:"re_password" binding:"required" msg:"请再次输入新密码"`
}

// UserUpdatePwdView 用户密码修改
// @Tags 用户管理
// @summary 用户密码修改
// @Description 用户密码修改
// @Param cr body UserUpdatePwdRequest true "用户密码修改"
// @Param token header string  true "token"
// @Router /user_update_pwd [put]
// @Produce json
// @success 200 {object} res.Response
func (UserApi) UserUpdatePwdView(c *gin.Context) {

	//	绑定参数
	var cr UserUpdatePwdRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//	查找用户
	var user models.UserModel
	err := global.DB.Find(&user, "id = ?", cr.UserID).Error
	if err != nil {
		res.FailWithMassage("该用户id未找到", c)
		return
	}

	//确认密码是否正确
	right := service.Service.UserService.CheckPwd(cr.UserID, cr.Password)
	if !right {
		res.FailWithMassage("用户密码错误", c)
		return
	}

	//确认新密码是否一致
	if cr.NewPassword != cr.RePassword {
		res.FailWithMassage("新密码两次输入不一致", c)
		return
	}

	//	用户密码修改
	passwordHash, err := utils.PasswordHash(cr.NewPassword)
	if err != nil {
		res.FailWithMassage("密码加密错误", c)
		return
	}
	err = global.DB.Model(&user).Update("password", passwordHash).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("修改用户密码出错", c)
		return
	}

	res.OkWithMassage("修改用户密码成功", c)

}
