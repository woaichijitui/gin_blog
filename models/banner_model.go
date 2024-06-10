package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`                 //图片路径
	Hash string `json:"hash"`                 //图片的hash值，用于判断重复的图片
	Name string `json:"name" gorm:"size:38"'` //图片名称
}
