package models

import "time"

// MODEL 不需要gorm.MODEL 中的删除时间
type MODEL struct {
	ID        uint      `json:"id" gorm:"primaryKey"` //主键id
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"-"`
}

// RemoveRequest 接收要删除的id
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

// PageInfo 接收前端的分页需求
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"` //每页的页数
	Sort  string `form:"sort"`
}
