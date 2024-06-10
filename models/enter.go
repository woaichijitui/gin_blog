package models

import "time"

type MODEL struct {
	ID        uint      `json:"id" gorm:"primaryKey"` //主键id
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"-"`
}
