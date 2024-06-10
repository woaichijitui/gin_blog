package models

import (
	"gvb_server/models/ctype"
)

type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"`                                                         //昵称
	UserName       string           `gorm:"size36" json:"user_name"`                                                          //用户名
	Password       string           `gorm:"size:128" json:"password"`                                                         //密码
	Avatar         string           `gorm:"size:246" json:"avatar"`                                                           //头像id
	Email          string           `gorm:"size:128" json:"email"`                                                            //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                               //手机
	Addr           string           `gorm:"size:64" json:"addr"`                                                              //地址
	Token          string           `gorm:"size:64" json:"token"`                                                             //其他平台唯一id
	IP             string           `gorm:"size:20" json:"ip"`                                                                //ip地址
	Role           ctype.Role       `gorm:"size:4,default:1" json:"role"`                                                     //用户权限
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                              //用户登录方式
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                       //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` //收藏的文章列表
}
