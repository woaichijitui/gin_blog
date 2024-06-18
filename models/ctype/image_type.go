package ctype

import "encoding/json"

type ImageType int

const (
	Local  ImageType = 1 //本地
	Aliyun ImageType = 2 //阿里云
)

// 序列化时调用这个函数
func (i ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())

}

// 角色映射
// role：角色
func (i ImageType) String() string {
	var str string
	switch i {
	case Local:
		str = "本地"
	case Aliyun:
		str = "阿里云"
	}
	return str
}
