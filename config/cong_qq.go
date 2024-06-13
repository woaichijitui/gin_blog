package config

import "fmt"

type QQ struct {
	AppID    string `yaml:"app_id" json:"app_id"`
	Key      string `yaml:"key" json:"key"`
	Redirect string `yaml:"redirect" json:"redirect"`
}

func (q *QQ) GetQQImage() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
