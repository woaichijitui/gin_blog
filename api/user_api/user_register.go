package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/utils"
	"gvb_server/utils/email"
)

type RegisterRequest struct {
	UserName   string  `json:"user_name,omitempty"  binding:"required" msg:"请正确输入用户名"`
	NickName   string  `json:"nick_name,omitempty" binding:"required" msg:"请正确输入昵称"`
	Password   string  `json:"password,omitempty" binding:"required" msg:"请正确输入密码"`
	RePassword string  `json:"re_password,omitempty" binding:"required" msg:"请正确输入密码"`
	Email      string  `json:"email,omitempty" binding:"required,email" msg:"请正确输入邮箱"`
	Code       *string `json:"code"`
}

// UserRegisterView 邮箱或者用户名注册
// @Tags 用户管理
// @summary 邮箱或者用户名注册
// @Description 邮箱或者用户名注册
// @Param cr body RegisterRequest true "注册信息 "
// @Router /user_register [post]
// @Produce json
// @success 200 {object} res.Response
func (l UserApi) UserRegisterView(c *gin.Context) {
	//	接收参数
	var cr RegisterRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//判断用户名是否存在
	var userModel models.UserModel
	row := global.DB.Find(&userModel, "user_name = ?", cr.UserName).RowsAffected
	if row > 0 {
		//	用户存在
		global.Log.Error("用户已经存在，请重新输入")
		return
	}
	//校验两次密码
	if cr.Password != cr.RePassword {
		//	密码不一致
		global.Log.Error("两次密码不一致")
		return
	}
	//hash加密密码
	hashPwd, err := utils.PasswordHash(cr.Password)
	if err != nil {
		global.Log.Error("加密密码失败：", err)
		return
	}

	//验证邮箱
	// 初始化session对象
	session := sessions.Default(c)
	//	第一次发送邮箱验证码
	if cr.Code == nil {

		//随机验证码
		code := utils.Code()
		//	发送验证码
		err = email.NewCode().SendEmail(cr.Email, fmt.Sprintf("你的验证码 %s", code))
		if err != nil {
			global.Log.Error(err)
			res.FailWithMassage("发送邮箱失败", c)
			return
		}
		//	将code存入session
		session.Set("mail_code", code)
		session.Set("email", cr.Email)
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
	if cr.Email != session.Get("email") {
		res.OkWithMassage("邮箱不一致", c)
		return
	}

	//头像
	//默认头像（地址
	avatar := "uploads/images/default.jpg"

	//	入库
	err = global.DB.Create(&models.UserModel{
		NickName:   cr.NickName,
		UserName:   cr.UserName,
		Password:   hashPwd,
		Avatar:     avatar,
		IP:         c.ClientIP(),
		Addr:       c.Request.URL.Path,
		Role:       ctype.PermissionUser,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error("创建用户失败：", err)
		return
	}
	global.Log.Infof("用户%s创建成功/n", cr.UserName)
	//	响应
	res.OkWithData("注册成功", c)
}
