package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	"gvb_server/utils/email"
)

type BindMailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"请输入邮箱"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

// UserBindMailView 用户绑定邮箱
// @Tags 用户管理
// @summary 用户绑定邮箱
// @Description 用户绑定邮箱
// @Param cr body BindMailRequest  true "用户绑定邮箱，第一次输入邮箱接收验证码，第二次输入验证码和密码（更新密码）"
// @Router /user_bind_email [post]
// @Produce json
// @success 200 {object} res.Response、
func (UserApi) UserBindMailView(c *gin.Context) {

	//	绑定
	var cr BindMailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	// 初始化session对象
	session := sessions.Default(c)
	//	第一次发送邮箱验证码
	if cr.Code == nil {

		//随机验证码
		code := utils.Code()
		//	发送验证码
		err = email.NewCode().SendEmail("1975611740@qq.com", fmt.Sprintf("你的验证码 %s", code))
		if err != nil {
			global.Log.Error(err)
			res.FailWithMassage("发送邮箱失败", c)
			return
		}
		//	将code存入session
		session.Set("mail_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMassage("session设置失败", c)
			return
		}
		res.OkWithMassage("验证码邮件已发送", c)
		return
	}
	//	第二次 验证邮箱验证码
	if *cr.Code != session.Get("mail_code") {
		res.OkWithMassage("验证码错误", c)
		return
	}

	//第一次邮箱 和 第二次邮箱一致性也要验证

	// 邮箱和密码更新
	_claims, exists := c.Get("claims")
	if !exists {
		global.Log.Error(err)
		res.OkWithMassage("获取claims错误", c)
		return
	}
	claims, ok := _claims.(*utils.MyCustomClaims)
	if !ok {
		global.Log.Error(err)
		res.OkWithMassage("获取claims错误", c)
		return
	}
	//	修改用户邮箱
	var userModel models.UserModel
	err = global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMassage("没有该用户", c)
		return
	}

	//	检查密码
	if len(cr.Password) < 4 {
		res.OkWithMassage("密码简单重新输入", c)
		return
	}
	hash, err := utils.PasswordHash(cr.Password)
	//	更新数据库
	err = global.DB.Model(&userModel).Updates(map[string]any{
		"password": hash,
		"email":    cr.Email,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.OkWithMassage("绑定邮箱失败", c)
		return
	}
	//	成功响应
	res.OkWithMassage("更新邮箱成功", c)

}
