package models

import "time"

// User2Collects 自定义第三章表 记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" `
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreateAt     time.Time
}
