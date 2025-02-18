package global

import (
	"github.com/sirupsen/logrus"
	"server/config"

	"gorm.io/gorm"
)

// 存放全局变量
var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
