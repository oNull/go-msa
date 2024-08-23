package global

import (
	"goods_srv/config"
	"gorm.io/gorm"
)

var (
	Db           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  = &config.NacosConfig{}
)
