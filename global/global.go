package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config *config.Config
	Log    *logrus.Logger
	Mysql  *gorm.DB
)
