package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	Config   *config.Config
	Log      *logrus.Logger
	DB       *gorm.DB
	MysqlLog logger.Interface
	Bucket   *oss.Bucket
	Redis    *redis.Client
)
