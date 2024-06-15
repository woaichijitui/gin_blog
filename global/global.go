package global

import (
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
)
