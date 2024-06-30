package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type UserUpdateRoleRequest struct {
	Role   ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"用户权限输入错误"`
	UserID uint       `json:"user_id" binding:"required" msg:"请输入用户id"`
}

// UserUpdateRole 用户权限修改
// @Tags 用户管理
// @summary 用户权限修改
// @Description 用户权限修改
// @Param cr body UserUpdateRoleRequest true "用户权限修改"
// @Router /user_update_role [put]
// @Produce json
// @success 200 {object} res.Response
func (UserApi) UserUpdateRole(c *gin.Context) {

	//	绑定参数
	var cr UserUpdateRoleRequest
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

	//	用户权限修改
	err = global.DB.Model(&user).Update("role", cr.Role).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("修改用户权限出错", c)
		return
	}

	res.OkWithMassage("修改用户权限成功", c)

}
