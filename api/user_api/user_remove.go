package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// UserRemoveView 批量删除用户
// @Tags 用户管理
// @summary 批量删除用户
// @Description 批量删除用户
// @Param cr body models.RemoveRequest true "要删除的用户id列表"
// @Router /user_delete [delete]
// @Produce json
// @success 200 {object} res.Response
func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	//	绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//	判断用户是否存在
	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMassage("用户不存在", c)
		return
	}

	//启动事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//todo: 关联表删除操作

		err := global.DB.Delete(userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		res.FailWithMassage("删除失败", c)
		return
	}

	//	成功响应
	res.OkWithMassage(fmt.Sprintf("成功删除 %d 个用户", count), c)
}
