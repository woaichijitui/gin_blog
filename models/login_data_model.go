package models

type LoginDataModel struct {
	MODEL
	UserID    uint      `  json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"`
	NickName  string    `gorm:"size:40" json:"nick_name"` //登录的ip
	Token     string    `gorm:"size:42" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` //登录的设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
