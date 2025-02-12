package models

import "gvb_server/models/ctype"

// MenuModel	菜单表
type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`                                                            //菜单标题
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en" `                                                        //英文菜单标题
	Slogan       string        `gorm:"size:64" json:"slogan" `                                                               //广告
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract" `                                                         //简介
	AbstractTime int           ` json:"abstract_time"`                                                                       //简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"-" ` //菜单图片列表
	BannerTime   int           `gorm:"" json:"banner_time" `                                                                 //菜单图片的切换时间
	Sort         int           `gorm:"size:10" json:"sort" `                                                                 //菜单的循序
}
