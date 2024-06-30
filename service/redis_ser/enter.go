package redis_ser

import (
	"context"
	"gvb_server/global"
	"time"
)

type RedisService struct {
}

// SetLogoutToken 放入已经注销的token
func (r RedisService) SetLogoutToken(token string, exp time.Duration) error {
	//	 将token和过期存入redis
	err := global.Redis.Set(context.Background(), "logout_"+token, "", exp).Err()
	return err
}

// CheckLogout 检查是否注销
func (r RedisService) CheckLogout(token string) bool {
	keys := global.Redis.Keys(context.Background(), "logout_*").Val()
	for _, key := range keys {
		if "logout_"+token == key {
			return true
		}
	}
	return false
}