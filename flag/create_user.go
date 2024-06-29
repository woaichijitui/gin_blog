package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
)

// CreateUser 命令行创建账号
func CreateUser(permission string) {

	//创建用户逻辑
	var (
		UserName   string
		NickName   string
		Password   string
		rePassword string
		Email      string
	)
	fmt.Print("请输入名字：")
	fmt.Scanln(&UserName)
	fmt.Print("请输入别名：")
	fmt.Scanln(&NickName)
	fmt.Print("请输入密码：")
	fmt.Scanln(&Password)
	fmt.Print("请输入再次输入：")
	fmt.Scanln(&rePassword)
	fmt.Print("请输入邮箱：")
	fmt.Scanln(&Email)

	//判断用户名是否存在
	var userModel models.UserModel
	row := global.DB.Find(&userModel, "user_name = ?", UserName).RowsAffected
	if row > 0 {
		//	用户存在
		global.Log.Error("用户已经存在，请重新输入")
		return
	}
	//校验两次密码
	if Password != rePassword {
		//	密码不一致
		global.Log.Error("两次密码不一致")
		return
	}
	//hash加密密码
	hashPwd, err := utils.PasswordHash(Password)
	if err != nil {
		global.Log.Error("加密密码失败：", err)
		return
	}

	//用户
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}
	//头像
	//默认头像（地址
	avatar := "uploads/images/default.jpg"

	//	入库
	err = global.DB.Create(&models.UserModel{
		NickName:   NickName,
		UserName:   UserName,
		Password:   hashPwd,
		Email:      Email,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		Role:       role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error("创建用户失败：", err)
		return
	}
	global.Log.Infof("用户%s创建成功/n", UserName)
}
