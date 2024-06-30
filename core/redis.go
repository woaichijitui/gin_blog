package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gvb_server/global"
	"time"
)

func ConnectRedisDB() *redis.Client {
	rc := global.Config.Redis
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr(),
		Password: rc.Password,
		DB:       rc.DB,
		PoolSize: rc.PoolSize,
	})

	//	context 超时控制等
	c, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//?
	_, err := rdb.Ping(c).Result()
	if err != nil {
		global.Log.Error(err)
		return nil
	}

	return rdb
}
