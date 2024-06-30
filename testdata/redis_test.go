package testdata

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestConnectRedis(t *testing.T) {

	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "120.27.192.46:6379",
		Password: "1234",
		DB:       0,
		PoolSize: 8,
	})
	//	context 超时控制等
	c, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	cmd, err := rdb.Ping(c).Result()
	if err != nil {

		return
	}
	fmt.Println(cmd)

	rdb.Set(c, "name", "1234", 8*time.Second)

	keys := rdb.Keys(c, "*")
	result, err := keys.Result()
	fmt.Println(result, err)

	name := rdb.Get(c, "name")
	fmt.Println(name)

}
