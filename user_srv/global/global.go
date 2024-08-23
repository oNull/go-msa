package global

import (
	"gorm.io/gorm"
	"user_srv/config"
)

var (
	Db           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  = &config.NacosConfig{}
)
