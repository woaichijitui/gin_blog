package service

import (
	"gvb_server/service/redis_ser"
	"gvb_server/service/user_ser"
)

type _Service struct {
	RedisService redis_ser.RedisService
	UserService  user_ser.UserService
}

var Service = new(_Service)
