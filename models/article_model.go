package models

import "gvb_server/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `json:"title" gorm:"size:32"`                           //文章标题
	Abstract      string         `json:"abstract"`                                       //文章简介
	Content       string         `json:"content"`                                        //文章内容
	LookCount     int            `json:"look_count"`                                     //浏览量
	CommentCount  int            `json:"comment_count"`                                  //评论量
	DiggCount     int            `json:"digg_count"`                                     //点赞量
	TagModels     []TagModel     `json:"tag_models" gorm:"many2many:article_tag_models"` //文章标签
	CommentModels []CommentModel `json:"-" gorm:"foreignKey:ArticleID"`                  //文章的评论列表
	UserModel     UserModel      `json:"-" gorm:"foreignKey:UserID"`                     //文章作者
	UserID        uint           `json:"user_id"`                                        //作者id
	Category      string         `json:"category" gorm:"size:20"`                        //文章分类
	Source        string         `json:"source"`                                         //文章来源
	Link          string         `json:"link"`                                           //原文连接
	Banner        BannerModel    `json:"-" gorm:"foreignKey:BannerID"`                   //文章封面
	BannerID      uint           `json:"banner_id"`                                      //文章封面ID
	NickName      string         `json:"nick_name" gorm:"size:42"`                       //发布文章的用户昵称
	BannerPath    string         `json:"banner_path"`                                    //文章的封面
	Tags          ctype.Array    `json:"tags" gorm:"size:64;type:string"`                //文章标签
}
